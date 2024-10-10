package middlewares

import (
	"go.opentelemetry.io/otel/metric"
	"gocourse20/pkg/telementry/meter"
	"log"
	"net/http"
	"time"
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

		meter.IncCounter(r.Context(), meter.TotalRequests)

		userName := r.Header.Get(`X-Username`)
		attr := meter.UserKey.String(userName)
		meter.UpdateGauge(r.Context(), meter.ConcurrentConnections, 1, metric.WithAttributes(attr))

		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
