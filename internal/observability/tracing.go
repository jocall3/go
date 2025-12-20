```go
// Copyright (c) 2023-2024 The Corredor Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package observability

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.24.0"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Config holds the configuration for the tracing setup.
type Config struct {
	Enabled     bool
	ServiceName string
	Endpoint    string // e.g., "otel-collector:4317"
	SampleRatio float64
	Insecure    bool // Use insecure gRPC connection for the OTLP exporter
}

// InitTracerProvider initializes and registers a new OpenTelemetry TracerProvider.
// It is responsible for creating the entire pipeline:
// Resource -> Exporter -> Span Processor -> Tracer Provider.
// It returns a shutdown function that should be deferred in the main function
// to ensure all buffered traces are flushed before the application exits.
func InitTracerProvider(cfg Config) (func(context.Context), error) {
	if !cfg.Enabled {
		log.Println("Tracing is disabled.")
		// Return a no-op shutdown function if tracing is not enabled.
		return func(context.Context) {}, nil
	}

	if cfg.ServiceName == "" {
		return nil, fmt.Errorf("service name must be provided for tracing")
	}
	if cfg.Endpoint == "" {
		return nil, fmt.Errorf("otlp endpoint must be provided for tracing")
	}

	log.Printf("Initializing tracing for service '%s' with endpoint '%s'", cfg.ServiceName, cfg.Endpoint)

	// Create a new resource with service-specific attributes.
	// A resource is a representation of the entity producing telemetry, which helps
	// to identify and query traces belonging to this specific service.
	res, err := resource.New(
		context.Background(),
		resource.WithAttributes(
			semconv.ServiceName(cfg.ServiceName),
			semconv.ServiceVersion("0.1.0"), // This could be made dynamic (e.g., from build info)
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create OTel resource: %w", err)
	}

	// Set up a connection to the OTLP collector.
	// In a production environment, this should use TLS credentials for security.
	connOpts := []grpc.DialOption{grpc.WithBlock()}
	if cfg.Insecure {
		connOpts = append(connOpts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	} else {
		// TODO: Implement TLS credentials for production environments.
		// For now, we log a warning and proceed with an insecure connection if not explicitly set.
		log.Println("WARNING: Using insecure gRPC connection for tracing. This is not recommended for production.")
		connOpts = append(connOpts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	// Use a context with a timeout for the initial connection to avoid blocking indefinitely.
	connCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(connCtx, cfg.Endpoint, connOpts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create gRPC connection to OTLP collector at %s: %w", cfg.Endpoint, err)
	}

	// Set up a trace exporter that sends data to the collector via gRPC.
	traceExporter, err := otlptracegrpc.New(context.Background(), otlptracegrpc.WithGRPCConn(conn))
	if err != nil {
		return nil, fmt.Errorf("failed to create OTLP trace exporter: %w", err)
	}

	// Configure the sampler. For a financial system, sampling everything (1.0) is
	// desirable for maximum auditability and debugging. This can be tuned down
	// for performance in high-volume, non-critical environments.
	var sampler sdktrace.Sampler
	if cfg.SampleRatio >= 1.0 {
		sampler = sdktrace.AlwaysSample()
	} else if cfg.SampleRatio <= 0.0 {
		sampler = sdktrace.NeverSample()
	} else {
		sampler = sdktrace.TraceIDRatioBased(cfg.SampleRatio)
	}

	// The BatchSpanProcessor is a standard processor that batches spans before exporting.
	// This is more efficient than sending each span individually.
	bsp := sdktrace.NewBatchSpanProcessor(traceExporter)

	// The TracerProvider is the core of the OpenTelemetry SDK.
	// It's responsible for creating Tracers and is configured with the resource,
	// sampler, and span processor.
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sampler),
		sdktrace.WithResource(res),
		sdktrace.WithSpanProcessor(bsp),
	)

	// Set the global TracerProvider. This allows any part of the application
	// to get a Tracer instance with otel.Tracer().
	otel.SetTracerProvider(tp)

	// Set the global Propagator. This is crucial for propagating trace context
	// across service boundaries (e.g., in HTTP headers or gRPC metadata).
	// W3C Trace Context is the standard for context propagation.
	// Baggage allows for carrying arbitrary key-value pairs along with the trace context.
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	))

	log.Println("Tracing initialized successfully.")

	// Return a shutdown function to be called on application exit.
	// This ensures that all buffered spans are sent to the collector and
	// the gRPC connection is closed gracefully.
	shutdown := func(ctx context.Context) {
		log.Println("Shutting down tracer provider...")
		shutdownCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()

		if err := tp.Shutdown(shutdownCtx); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
		if err := conn.Close(); err != nil {
			log.Printf("Error closing gRPC connection to collector: %v", err)
		}
		log.Println("Tracer provider shut down.")
	}

	return shutdown, nil
}

// GetTracer returns a named tracer from the global provider.
// It's a convenience wrapper around otel.Tracer.
// The 'name' should be the package or component name creating the spans,
// following OpenTelemetry instrumentation conventions.
func GetTracer(name string) otel.Tracer {
	return otel.Tracer(name)
}
### END_OF_FILE_COMPLETED ###
```