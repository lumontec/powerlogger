package powerlogger

import (
	"context"
	"log"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp"

	//	"go.opentelemetry.io/otel/exporters/stdout"
	"go.opentelemetry.io/otel/label"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/metric/controller/push"
	"go.opentelemetry.io/otel/sdk/metric/processor/basic"
	"go.opentelemetry.io/otel/sdk/metric/selector/simple"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/semconv"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
)

var plogger = &powerlogger{}

// Config powerlogger global configuration
type Config struct {
	ServiceName    string
	ServiceVersion string
	HostHostname   string
	CollectorAddr  string
	PusherPeriod   time.Duration
}

// PowerLogger glogbal object
type powerlogger struct {
	//	consolexp      *stdout.Exporter
	logger         *zap.Logger
	collectorexp   *otlp.Exporter
	pusher         *push.Controller
	tracerprovider *sdktrace.TracerProvider
	tracer         trace.Tracer
}

// Start initializes powerlogger
func Start(lc Config) context.Context {
	ctx := initTracer(lc.CollectorAddr, lc.ServiceName, lc.PusherPeriod)
	initLogger()
	return ctx
}

func initLogger() {
	logger, _ := newLoggerConfig().Build(zap.AddCaller(), zap.AddCallerSkip(1))
	defer logger.Sync()
	plogger.logger = logger
}

func newLoggerConfig() zap.Config {
	return zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         "json",
		EncoderConfig:    newLoggerEncoderConfig(),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}
}

func newLoggerEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:  "ts",
		LevelKey: "level",
		//		NameKey:        "logger",
		//		CallerKey:      "span",
		FunctionKey:    "span",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.EpochTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

//func Example_advancedConfiguration() {
//
//	logger := zap.New(core)
//	defer logger.Sync()
//	logger.Info("constructed a logger")
//}
//

func initTracer(collectorAddr string, serviceName string, pusherPeriod time.Duration) context.Context {

	ctx := context.Background()

	collectorexp, err := otlp.NewExporter(
		otlp.WithInsecure(),
		otlp.WithAddress(collectorAddr),
		otlp.WithGRPCDialOption(grpc.WithBlock()), // useful for testing
	)

	plogger.collectorexp = collectorexp

	handleErr(err, "failed to create collectorexp")

	//	consoleexp, err := stdout.NewExporter([]stdout.Option{
	//		stdout.WithQuantiles([]float64{0.5, 0.9, 0.99}),
	//		stdout.WithPrettyPrint(),
	//	}...)
	//
	//	plogger.consolexp = consoleexp

	//	handleErr(err, "failed to create consoleexp")

	res, err := resource.New(ctx,
		resource.WithAttributes(
			// the service name used to display traces in backends
			semconv.ServiceNameKey.String(serviceName),
		),
	)
	handleErr(err, "failed to create resource")

	collectorbsp := sdktrace.NewBatchSpanProcessor(plogger.collectorexp)
	//	consolebsp := sdktrace.NewBatchSpanProcessor(plogger.consolexp)

	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithConfig(sdktrace.Config{DefaultSampler: sdktrace.AlwaysSample()}),
		sdktrace.WithResource(res),
		sdktrace.WithSpanProcessor(collectorbsp),
		//		sdktrace.WithSpanProcessor(consolebsp),
	)

	plogger.tracerprovider = tracerProvider

	pusher := push.New(
		basic.New(
			simple.NewWithExactDistribution(),
			collectorexp,
		),
		collectorexp,
		push.WithPeriod(pusherPeriod),
	)

	// set global propagator to tracecontext (the default is no-op).
	otel.SetTextMapPropagator(propagation.TraceContext{})
	otel.SetTracerProvider(tracerProvider)
	otel.SetMeterProvider(pusher.MeterProvider())
	pusher.Start()

	plogger.tracer = otel.Tracer("global-tracer")

	return ctx

}

// Stop powerlogger
func Stop(ctx context.Context) {
	handleErr(plogger.tracerprovider.Shutdown(ctx), "failed to shutdown provider")
	handleErr(plogger.collectorexp.Shutdown(ctx), "failed to stop collector exporter")
	//	handleErr(plogger.consolexp.Shutdown(ctx), "failed to stop console exporter")
	plogger.pusher.Stop() // pushes any last exports to the receiver
}

// Span Generates child span and logger for context
func Span(ctx context.Context) context.Context {
	childCtx, _ := plogger.tracer.Start(ctx, "temp")
	return childCtx
}

// CloseSpan closes current span, assigns name, sets the status
func CloseSpan(ctx context.Context) {
	span := trace.SpanFromContext(ctx)
	// Skip the first 3 frames to pass CloseSpan
	callername := callerFrameName(3)
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
