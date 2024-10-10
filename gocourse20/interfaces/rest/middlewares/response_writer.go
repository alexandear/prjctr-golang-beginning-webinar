package middlewares

import (
	"context"
	"gocourse20/pkg/telementry/meter"
	"net/http"
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
		meter.IncCounter(rw.ctx, meter.TotalSuccessResponse)
	default:
		meter.IncCounter(rw.ctx, meter.TotalErrorsResponse)
	}

	rw.ResponseWriter.WriteHeader(code)
}

func (rw *responseWriter) Write(body []byte) (int, error) {
	rw.timeTracker.Flush(rw.ctx, meter.RequestDuration)
	meter.UpdateGauge(rw.ctx, meter.ConcurrentConnections, -1)

	return rw.ResponseWriter.Write(body)
}
