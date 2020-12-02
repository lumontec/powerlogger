package powerlogger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCallerFrameName(t *testing.T) {
	callerframe := subfunc1()
	assert.Equal(t, "powerlogger.subfunc1", callerframe, "callerframename should return subfunc1 caller name")
}

func BenchmarkCallerFrameName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		subfunc1()
	}
}

func subfunc1() string {
	return subfunc2()
}

func subfunc2() string {
	return callerFrameName(3)
}
