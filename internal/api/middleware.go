package api

import (
	"net/http"
	logger "phone-number-manager/internal/logging"

	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
)

// LoggingMiddleware adds structured request logs.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Log.Info("Incoming request", zap.String("method", r.Method), zap.String("url", r.URL.String()), zap.String("request_id", r.Header.Get("X-Request-ID")))
		next.ServeHTTP(w, r)
	})
}

// TracingMiddleware adds distributed tracing to requests.
func TracingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tracer := otel.Tracer("phone-number-manager")
		ctx, span := tracer.Start(r.Context(), r.URL.Path)
		defer span.End()
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
