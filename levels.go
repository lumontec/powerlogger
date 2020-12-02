package powerlogger

import (
	"context"
	"go.opentelemetry.io/otel/trace"
)

// Info emits a log with Info level
func Info(ctx context.Context, msg string, labels ...Label) {
	span := trace.SpanFromContext(ctx)
	otelLabels := OtelLabels(labels...)
	zapLabels := ZapLabels(labels...)
	span.AddEvent(msg, trace.WithAttributes(otelLabels...))
	plogger.logger.Info(msg, zapLabels...)
}

// Debug emits a log with Debug level
func Debug(ctx context.Context, msg string, labels ...Label) {
	span := trace.SpanFromContext(ctx)
	otelLabels := OtelLabels(labels...)
	zapLabels := ZapLabels(labels...)
	span.AddEvent(msg, trace.WithAttributes(otelLabels...))
	plogger.logger.Debug(msg, zapLabels...)
}

//  Error emits a log with Error level
func Error(ctx context.Context, msg string, labels ...Label) {
	span := trace.SpanFromContext(ctx)
	otelLabels := OtelLabels(labels...)
	zapLabels := ZapLabels(labels...)
	span.AddEvent(msg, trace.WithAttributes(otelLabels...))
	plogger.logger.Error(msg, zapLabels...)
}
