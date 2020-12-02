package powerlogger

import (
	"context"
	"go.opentelemetry.io/otel/label"
	"go.opentelemetry.io/otel/trace"
)

// Info emits a log with Info level
func Info(ctx context.Context, msg string, labels ...label.KeyValue) {
	span := trace.SpanFromContext(ctx)
	span.AddEvent(msg, trace.WithAttributes(labels...))
	plogger.logger.Info(msg)
}

// Debug emits a log with Debug level
func Debug(ctx context.Context, msg string, labels ...label.KeyValue) {
	span := trace.SpanFromContext(ctx)
	span.AddEvent(msg, trace.WithAttributes(labels...))
	plogger.logger.Debug(msg)
}

//  Error emits a log with Error level
func Error(ctx context.Context, msg string, labels ...label.KeyValue) {
	span := trace.SpanFromContext(ctx)
	span.AddEvent(msg, trace.WithAttributes(labels...))
	plogger.logger.Error(msg)
}
