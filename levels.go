package powerlogger

import (
	"context"

	"go.opentelemetry.io/otel/baggage"
	"go.opentelemetry.io/otel/label"
	"go.opentelemetry.io/otel/trace"
)

// Info emits a log with Info level
func Info(ctx context.Context, msg string, labels ...Label) {
	span := trace.SpanFromContext(ctx)
	otelLabels := OtelLabels(labels...)
	otelLabels = append(otelLabels, label.String("level", "info"))
	span.AddEvent(msg, trace.WithAttributes(otelLabels...))
	zapLabels := ZapLabels(labels...)
	injLabelSet := baggage.Set(ctx)
	injLabArr := injLabelSet.ToSlice()
	for _, injLab := range injLabArr {
		zapLabels = append(zapLabels, ParseOtelLabel(injLab).ZapLabel())
	}
	plogger.logger.Info(msg, zapLabels...)
}

// Debug emits a log with Debug level
func Debug(ctx context.Context, msg string, labels ...Label) {
	span := trace.SpanFromContext(ctx)
	otelLabels := OtelLabels(labels...)
	otelLabels = append(otelLabels, label.String("level", "debug"))
	span.AddEvent(msg, trace.WithAttributes(otelLabels...))
	zapLabels := ZapLabels(labels...)
	plogger.logger.Debug(msg, zapLabels...)
}

//  Error emits a log with Error level
func Error(ctx context.Context, msg string, labels ...Label) {
	span := trace.SpanFromContext(ctx)
	otelLabels := OtelLabels(labels...)
	otelLabels = append(otelLabels, label.String("level", "error"))
	span.AddEvent(msg, trace.WithAttributes(otelLabels...))
	zapLabels := ZapLabels(labels...)
	plogger.logger.Error(msg, zapLabels...)
}
