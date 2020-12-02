package powerlogger

import (
	"runtime"
)

func callerFrameName(skip int) string {
	counter := make([]uintptr, 4)
	// Skip first skip frames in the stack
	runtime.Callers(skip, counter)
	frames := runtime.CallersFrames(counter)
	// Pull first frame
	frame, _ := frames.Next()
	return frame.Function
}
