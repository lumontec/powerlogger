package powerlogger

import (
	"runtime"
)

func callerFrameName() string {
	counter := make([]uintptr, 4)
	// Skip first 2 frames in the stack
	runtime.Callers(2, counter)
	frames := runtime.CallersFrames(counter)
	// Pull first frame
	frame, _ := frames.Next()
	return frame.Function
}
