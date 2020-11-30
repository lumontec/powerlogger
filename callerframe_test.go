package powerlogger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCallerFrameName(t *testing.T) {
	callerframe := subfunc()
	assert.Equal(t, "powerlogger.subfunc", callerframe, "callerframename should return subfunc caller name")
}

func BenchmarkCallerFrameName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		subfunc()
	}
}

func subfunc() string {
	return callerFrameName()
}
