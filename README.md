# **POWERLOGGER** 
Powerlogger is the logger that you expect, enhanced with software telemetry superpowers

## Rationale 
Lets face it, implementing telemetry is getting complex and verbose, and not everybody has the time or the resources to properly and fully instrument new or old codebases. 
Part of this complexity is due to the vastity of configurations and deployment options allowed by the libraries. 

Powerlogger comes to help you avoiding this complexity by providing an opinionated implementation that runs most of the stuff under the hood.
The core idea of the project is allowing everyone to refactor his old logger with powerlogger, gaining telemetric functionalities at little to no cost, while avoiding polluting the codebase with excessive amount of infrasctructure related stuff.

## Functionality 
Powerlogger exposes a simple minimalistic logging api, implementing most of the opinionated code telemetry under the hood
- json logging as a mission 
- implements a singleton global object, no need to pass you logger around anymore 
- automatically generates span names from function callers (efficiently gathered from the caller frame) 
- tracks spans through context propagation
- tees information to both console and opentelemetry-collector by default 

## **Example API** 

*Import powerlogger:* 
```go
import ( plog "powerlogger")
```

*Initialize powerlogger:* 
```go
ctx := plog.Start(plog.Config{
	ServiceName:    "test_service",
	ServiceVersion: "1.0.0",
	HostHostname:   "localhost",
	CollectorAddr:  "localhost:55680",
	PusherPeriod:   7 * time.Second,
})
```

*Generate new span for func sub1:* 
```go
sub1(plog.Span(ctx), "arg1")
```

*Close span for function:*
```go
defer plog.CloseSpan(ctx)
```

*Log event with some context (add log to span):*
```go
plog.Debug(ctx, "inside function sub1", plog.Bool("customkey3", true))
```

*Inject key value context to downstream spans:*
```go
plog.Inject(ctx, plog.Bool("injectedkey3", false))
```

*Full example (./example/main.go):*
```go
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
	defer plog.CloseSpan(ctx, "sub1")

	plog.Debug(ctx, "inside function sub1", plog.Bool("customkey3", true))
	plog.Inject(ctx, plog.Bool("injectedkey1", false))

	sub2(plog.Span(ctx), "1arg", false)
}

func sub2(ctx context.Context, arg1 string, arg2 bool) {
	defer plog.CloseSpan(ctx, "sub2")

	plog.Debug(ctx, "inside function sub2", plog.Bool("customkey4", true))

}
```



