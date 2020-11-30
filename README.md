# **POWERLOGGER** 
Powerlogger is the logger that you expect, enhanced with software telemetry superpowers

## Rationale 
Lets face it, implementing telemetry is getting complex and verbose, and not everybody has the time or the resources to properly and fully instrument new or old codebases. Part of this complexity is due to the vastity of configurations and deployment options allowed by the libraries. Powerlogger comes to help you avoiding this complexity by providing an opinionated implementation that runs most of the stuff under the hood.
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
```
import ( plog "powerlogger")
```

*Initialize powerlogger:* 
```
ctx := plog.Start(plog.Config{
	ServiceName:    "test_service",
	ServiceVersion: "1.0.0",
	HostHostname:   "localhost",
	CollectorAddr:  "localhost:55680",
	PusherPeriod:   7 * time.Second,
})
```

*Generate new span for func sub1:* 
```
sub1(plog.Span(ctx), 1)
```

*Close span for function:*
```
defer plog.CloseSpan(ctx)
```

*Log event with some context (add log to span):*
```
plog.Debug(ctx, "inside function sub1", plog.Bool("customkey3", true))
```

*Inject key value context to downstream spans:*
```
plog.Inject(ctx, plog.Bool("injectedkey3", false))
```

