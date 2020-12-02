# Basic powerlogger usage example 

This directory contains a simple example for basic usage of powerlogger 

## Content
This example is composed from three components:
- *Our application*(main.go): will log to console and trace to the collector 
- *OpenTelemetry collector(./infra/docker-compose)*: will ingest spans and metrics and send them to Prometheus and Jaeger
- *Prometheus(./infra/docker-compose)*: processes and expose metric information
- *Jaeger(./infra/docker-compose)*: processes and expose tracing information 

## Run
1- ```cd ./infra && docker-compose up -d```
2- ```go run main.go```

Apache License 2.0
