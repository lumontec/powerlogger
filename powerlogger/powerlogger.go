package powerlogger

import (
	"context"

	"go.opentelemetry.io/otel/label"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	// "go.uber.org/zap"
	// "go.uber.org/zap/zapcore"
)

var glogger = &Logger{}

// Config configuration
type Config struct {
	Ciao bool
}

// Logger object
type Logger struct {
	zlog *zap.Logger
}

// Label data type
type Label interface{}

// Start initializes the logger
func Start(lc Config) {
	glogger = &Logger{}
}

// Span Generates child span and logger for context
func Span(ctx context.Context, attr ...zapcore.Field) context.Context {
	return ctx
}

// CloseSpan closes current span, assigns name, sets the status
func CloseSpan(ctx context.Context) {
}

// Inject injects custom key values inside context
func Inject(ctx context.Context, labels ...Label) {
}

// Debug logs with Debug level
func Debug(ctx context.Context, msg string, labels ...Label) {
	// glogger.zlog.Debug(msg, attr...)
}

// Bool attach Bool label
func Bool(key string, val bool) label.KeyValue {
	return label.Bool(key, val)
}
