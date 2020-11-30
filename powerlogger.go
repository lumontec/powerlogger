package powerlogger

import (
	"context"
	"log"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp"
	"go.opentelemetry.io/otel/exporters/stdout"
	"go.opentelemetry.io/otel/label"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/metric/controller/push"
	"go.opentelemetry.io/otel/sdk/metric/processor/basic"
	"go.opentelemetry.io/otel/sdk/metric/selector/simple"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/semconv"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	// "go.uber.org/zap"
	// "go.uber.org/zap/zapcore"
)

var glogger = &powerlogger{}

// Config configuration
type Config struct {
	ServiceName    string
	ServiceVersion string
	HostHostname   string
	CollectorAddr  string
	PusherPeriod   time.Duration
}

// Logger object
type powerlogger struct {
	consolexp      *stdout.Exporter
	collectorexp   *otlp.Exporter
	pusher         *push.Controller
	tracerprovider *sdktrace.TracerProvider
	tracer         trace.Tracer
}

// Label data type
type Label interface{}

// Start initializes powerlogger
func Start(lc Config) context.Context {

	ctx := context.Background()

	collectorexp, err := otlp.NewExporter(
		otlp.WithInsecure(),
		otlp.WithAddress(lc.CollectorAddr),
		otlp.WithGRPCDialOption(grpc.WithBlock()), // useful for testing
	)

	glogger.collectorexp = collectorexp

	handleErr(err, "failed to create collectorexp")

	consoleexp, err := stdout.NewExporter([]stdout.Option{
		stdout.WithQuantiles([]float64{0.5, 0.9, 0.99}),
		stdout.WithPrettyPrint(),
	}...)

	glogger.consolexp = consoleexp

	handleErr(err, "failed to create consoleexp")

	res, err := resource.New(ctx,
		resource.WithAttributes(
			// the service name used to display traces in backends
			semconv.ServiceNameKey.String(lc.ServiceName),
		),
	)
	handleErr(err, "failed to create resource")

	collectorbsp := sdktrace.NewBatchSpanProcessor(glogger.collectorexp)
	consolebsp := sdktrace.NewBatchSpanProcessor(glogger.consolexp)

	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithConfig(sdktrace.Config{DefaultSampler: sdktrace.AlwaysSample()}),
		sdktrace.WithResource(res),
		sdktrace.WithSpanProcessor(collectorbsp),
		sdktrace.WithSpanProcessor(consolebsp),
	)

	glogger.tracerprovider = tracerProvider

	pusher := push.New(
		basic.New(
			simple.NewWithExactDistribution(),
			collectorexp,
		),
		collectorexp,
		push.WithPeriod(lc.PusherPeriod),
	)

	// set global propagator to tracecontext (the default is no-op).
	otel.SetTextMapPropagator(propagation.TraceContext{})
	otel.SetTracerProvider(tracerProvider)
	otel.SetMeterProvider(pusher.MeterProvider())
	pusher.Start()

	glogger.tracer = otel.Tracer("global-tracer")

	return ctx
}

// Stop powerlogger
func Stop(ctx context.Context) {
	handleErr(glogger.tracerprovider.Shutdown(ctx), "failed to shutdown provider")
	handleErr(glogger.collectorexp.Shutdown(ctx), "failed to stop collector exporter")
	handleErr(glogger.consolexp.Shutdown(ctx), "failed to stop console exporter")
	glogger.pusher.Stop() // pushes any last exports to the receiver
}

// Span Generates child span and logger for context
func Span(ctx context.Context) context.Context {
	childCtx, _ := glogger.tracer.Start(ctx, "temp")
	return childCtx
}

// CloseSpan closes current span, assigns name, sets the status
func CloseSpan(ctx context.Context, name string) {
	span := trace.SpanFromContext(ctx)
	callername := callerFrameName()
	span.SetName(callername)
	span.End()
}

// Inject injects custom key values inside context
func Inject(ctx context.Context, labels ...label.KeyValue) {
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(labels...)
}

func handleErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s: %v", message, err)
	}
}
