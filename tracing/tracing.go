package tracing

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/trace"
)

// Init sets up tracing and returns a shutdown function.
func Init(serviceName string) func(context.Context) error {
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint())
	if err != nil {
		panic(err)
	}
	tp := trace.NewTracerProvider(
		trace.WithBatcher(exp),
	)
	otel.SetTracerProvider(tp)
	return tp.Shutdown
}