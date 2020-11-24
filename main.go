package main

import (
	"context"
	"fmt"
	plog "main/powerlogger"
)

func main() {

	fmt.Println("Initializing the logger...")

	plog.Start(plog.Config{})

	ctx := context.Background()

	plog.Debug(ctx, "main function initialized", plog.Bool("customkey1", true), plog.Bool("customkey2", true))
	plog.Inject(ctx, plog.Bool("injectedkey1", false))

	sub1(plog.Span(ctx), 1)
}

func sub1(ctx context.Context, arg1 int) {
	defer plog.CloseSpan(ctx)

	plog.Debug(ctx, "inside function sub1", plog.Bool("customkey3", true))
	plog.Inject(ctx, plog.Bool("injectedkey2", false))

	sub2(plog.Span(ctx), "1arg", false)
}

func sub2(ctx context.Context, arg1 string, arg2 bool) {
	defer plog.CloseSpan(ctx)

	plog.Debug(ctx, "inside function sub2", plog.Bool("customkey4", true))
}
