package powerlogger

import (
	"go.opentelemetry.io/otel/label"
	"go.uber.org/zap"
)

type LabelType int

const (
	// INVALID is used for a Value with no value set.
	INVALID LabelType = iota
	BOOL
	INT32
	INT64
	UINT32
	UINT64
	FLOAT32
	FLOAT64
	STRING
	ARRAY
)

// Label struct for context information
type Label interface {
	OtelLabel() label.KeyValue
	ZapLabel() zap.Field
}

func OtelLabels(labels ...Label) []label.KeyValue {
	otlabels := []label.KeyValue{}
	for _, label := range labels {
		otlabels = append(otlabels, label.OtelLabel())
	}
	return otlabels
}

func ZapLabels(labels ...Label) []zap.Field {
	zaplabels := []zap.Field{}
	for _, label := range labels {
		zaplabels = append(zaplabels, label.ZapLabel())
	}
	return zaplabels
}

func ParseOtelLabel(label label.KeyValue) Label {
	switch LabelType(label.Value.Type()) {
	case BOOL:
		return Bool(string(label.Key), label.Value.AsBool())
	case INT32:
		return Int32(string(label.Key), label.Value.AsInt32())
	case INT64:
		return Int64(string(label.Key), label.Value.AsInt64())
	// case UINT32:
	// 	return Uint64(string(label.Key), label.Value.AsInt64())
	// case UINT64:
	case FLOAT32:
		return Float32(string(label.Key), label.Value.AsFloat32())
	case FLOAT64:
		return Float64(string(label.Key), label.Value.AsFloat64())
	case STRING:
		return String(string(label.Key), label.Value.AsString())
	case ARRAY:
	}
	return nil
}

type BoolLabel struct {
	Type LabelType
	Key  string
	Val  bool
}

type IntLabel struct {
	Type LabelType
	Key  string
	Val  int
}

type Int32Label struct {
	Type LabelType
	Key  string
	Val  int32
}

type Int64Label struct {
	Type LabelType
	Key  string
	Val  int64
}

type Float32Label struct {
	Type LabelType
	Key  string
	Val  float32
}

type Float64Label struct {
	Type LabelType
	Key  string
	Val  float64
}

type StringLabel struct {
	Type LabelType
	Key  string
	Val  string
}

// Bool attach Bool label
func Bool(key string, val bool) *BoolLabel {
	return &BoolLabel{
		Type: BOOL,
		Key:  key,
		Val:  val,
	}
}

// Int32 attach Int32 label
func Int32(key string, val int32) *Int32Label {
	return &Int32Label{
		Type: INT32,
		Key:  key,
		Val:  val,
	}
}

// Int64 attach Int64 label
func Int64(key string, val int64) *Int64Label {
	return &Int64Label{
		Type: INT64,
		Key:  key,
		Val:  val,
	}
}

// Float32 attach Float32 label
func Float32(key string, val float32) *Float32Label {
	return &Float32Label{
		Type: FLOAT32,
		Key:  key,
		Val:  val,
	}
}

// Float64 attach Float64 label
func Float64(key string, val float64) *Float64Label {
	return &Float64Label{
		Type: FLOAT64,
		Key:  key,
		Val:  val,
	}
}

// String attach String label
func String(key string, val string) *StringLabel {
	return &StringLabel{
		Type: STRING,
		Key:  key,
		Val:  val,
	}
}

func (bl *BoolLabel) OtelLabel() label.KeyValue {
	return label.Bool(bl.Key, bl.Val)
}

func (bl *BoolLabel) ZapLabel() zap.Field {
	return zap.Bool(bl.Key, bl.Val)
}

func (il *IntLabel) OtelLabel() label.KeyValue {
	return label.Int(il.Key, il.Val)
}

func (il *IntLabel) ZapLabel() zap.Field {
	return zap.Int(il.Key, il.Val)
}

func (il32 *Int32Label) OtelLabel() label.KeyValue {
	return label.Int32(il32.Key, il32.Val)
}

func (il32 *Int32Label) ZapLabel() zap.Field {
	return zap.Int32(il32.Key, il32.Val)
}

func (il64 *Int64Label) OtelLabel() label.KeyValue {
	return label.Int64(il64.Key, il64.Val)
}

func (il64 *Int64Label) ZapLabel() zap.Field {
	return zap.Int64(il64.Key, il64.Val)
}

func (fl32 *Float32Label) OtelLabel() label.KeyValue {
	return label.Float32(fl32.Key, fl32.Val)
}

func (fl32 *Float32Label) ZapLabel() zap.Field {
	return zap.Float32(fl32.Key, fl32.Val)
}

func (fl64 *Float64Label) OtelLabel() label.KeyValue {
	return label.Float64(fl64.Key, fl64.Val)
}

func (fl64 *Float64Label) ZapLabel() zap.Field {
	return zap.Float64(fl64.Key, fl64.Val)
}

func (sl *StringLabel) OtelLabel() label.KeyValue {
	return label.String(sl.Key, sl.Val)
}

func (sl *StringLabel) ZapLabel() zap.Field {
	return zap.String(sl.Key, sl.Val)
}
