package main

import (
	"context"
	"fmt"
	"time"

	plog "powerlogger"
)

func main() {

	fmt.Println("Initializing the logger...")

	ctx := plog.Start(plog.Config{
		ServiceName:    "test_service",
		ServiceVersion: "1.0.0",
		HostHostname:   "localhost",
		CollectorAddr:  "localhost:55680",
		PusherPeriod:   7 * time.Second,
	})

	plog.Debug(ctx, "main function initialized", plog.Bool("customkey1", true), plog.Bool("customkey2", true))

	sub1(plog.Span(ctx), 1)
	time.Sleep(10 * time.Second)
}

func sub1(ctx context.Context, arg1 int) {
	defer plog.CloseSpan(ctx)

	plog.Debug(ctx, "inside function sub1", plog.Bool("customkey3", true))
	plog.Inject(ctx, plog.Bool("injectedkey1", false))

	sub2(plog.Span(ctx), "1arg", false)
}

func sub2(ctx context.Context, arg1 string, arg2 bool) {
	defer plog.CloseSpan(ctx)

	plog.Debug(ctx, "inside function sub2", plog.Bool("customkey4", true))

}
