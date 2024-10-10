package middleware

import (
	"context"
	"net/http"

	"gocourse20/internal/telemetry/meter"
)

type responseWriter struct {
	ctx context.Context
	http.ResponseWriter

	timeTracker meter.TimeTracker
}

func NewResponseWriter(ctx context.Context, w http.ResponseWriter, tt meter.TimeTracker) *responseWriter {
	return &responseWriter{ctx, w, tt}
}

func (rw *responseWriter) WriteHeader(code int) {
	switch code {
	case 200:
		meter.MustIncCounter(rw.ctx, meter.TotalSuccessResponse)
	default:
		meter.MustIncCounter(rw.ctx, meter.TotalErrorsResponse)
	}

	rw.ResponseWriter.WriteHeader(code)
}

func (rw *responseWriter) Write(body []byte) (int, error) {
	rw.timeTracker.MustFlush(rw.ctx, meter.RequestDuration)
	meter.MustUpdateGauge(rw.ctx, meter.ConcurrentConnections, -1)

	return rw.ResponseWriter.Write(body)
}
