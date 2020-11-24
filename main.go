package main

import (
	"context"
	"fmt"
	plog "powerlogger/logger"
	"runtime"
)

func main() {

	fmt.Println("Starting logger...")
	plog.Start(&plog.Config{})
	// stack := stacktrace.TakeStacktrace(1)
	// fmt.Printf("main stack: %v\n", stack)

	ctx := context.Background()

	plog.Debug(ctx, "Main function initialized")

	plog.Inject(ctx, plog.Bool("ciao", true))

	sub1(plog.Span(ctx))
}

func sub1(ctx context.Context) {
	defer plog.CloseSpan(ctx)

	plog.Debug(ctx, "Inside func sub1")
	sub2(plog.Span(ctx))
}

func sub2(ctx context.Context) {
	defer plog.CloseSpan(ctx)

	// stack := stacktrace.TakeStacktrace(0)
	// fmt.Printf("sub2 stack: %v\n", stack)
	// counter := make([]uintptr, 64)
	// runtime.Callers(0, counter)
	// frames := runtime.CallersFrames(counter[0:3])
	// fmt.Printf("sub2 pcs: %v\n", frames)

	buffer := make([]byte, 1000)
	runtime.Stack(buffer, false)
	fmt.Printf("%v", string(buffer))
}
