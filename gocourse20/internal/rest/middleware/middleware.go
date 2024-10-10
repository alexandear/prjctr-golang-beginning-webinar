package middleware

import (
	"log"
	"net/http"
	"time"

	"go.opentelemetry.io/otel/metric"

	"gocourse20/internal/telemetry/meter"
)

func Last(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		rw := NewResponseWriter(r.Context(), w, meter.TrackTimeBegin())
		next.ServeHTTP(rw, r)

		log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
	})
}

// Meter is a middleware that sets common response headers.
func Meter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		meter.MustIncCounter(r.Context(), meter.TotalRequests)

		userName := r.Header.Get("X-Username")
		attr := meter.UserKey.String(userName)
		meter.MustUpdateGauge(r.Context(), meter.ConcurrentConnections, 1, metric.WithAttributes(attr))

		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
