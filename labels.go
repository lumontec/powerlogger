package powerlogger

import (
	"go.opentelemetry.io/otel/label"
)

// Bool attach Bool label
func Bool(key string, val bool) label.KeyValue {
	return label.Bool(key, val)
}

// Int attach Int label
func Int(key string, val int) label.KeyValue {
	return label.Int(key, val)
}

// Int32 attach Int32 label
func Int32(key string, val int32) label.KeyValue {
	return label.Int32(key, val)
}

// Int64 attach Int64 label
func Int64(key string, val int64) label.KeyValue {
	return label.Int64(key, val)
}

// Float32 attach Float32 label
func Float32(key string, val float32) label.KeyValue {
	return label.Float32(key, val)
}

// Float64 attach Float64 label
func Float64(key string, val float64) label.KeyValue {
	return label.Float64(key, val)
}
