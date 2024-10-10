// Package meter wraps the prometheus instrumentation
package meter

import (
	"context"
	"go.opentelemetry.io/otel/attribute"
	api "go.opentelemetry.io/otel/metric"
	"log"
	"net/http"
	"time"
)

const packageName = `gocourse20`

type TimeTracker time.Time

const (
	ConcurrentConnections = iota
	RequestDuration
	TotalErrorsResponse
	TotalRequests
	TotalSuccessResponse
)

type baseMetric struct {
	name        string
	description string
}

type countMetric struct {
	baseMetric
	fn func(ctx context.Context, val float64, attrs ...api.AddOption)
}

type gaugeMetric struct {
	baseMetric
	fn func(ctx context.Context, val float64, attrs ...api.ObserveOption)
}

type histMetric struct {
	baseMetric
	fn func(ctx context.Context, val float64, attrs ...api.RecordOption)
}

var counters = map[uint8]*countMetric{
	TotalRequests: {
		baseMetric: baseMetric{
			name:        "request_sum_total",
			description: `Total number of requests processed`,
		},
	},
	TotalErrorsResponse: {
		baseMetric: baseMetric{
			name:        "response_error_sum_total",
			description: `Number of requests processed with an error`,
		},
	},
	TotalSuccessResponse: {
		baseMetric: baseMetric{
			name:        "response_success_sum_total",
			description: `Number of successfully processed requests`},
	},
}

var gauges = map[uint8]*gaugeMetric{
	ConcurrentConnections: {
		baseMetric: baseMetric{
			name:        "concurrent_connections",
			description: `The number of concurrent connections at the moment`,
		},
	},
}

var histograms = map[uint8]*histMetric{
	RequestDuration: {
		baseMetric: baseMetric{
			name:        "request_duration",
			description: `Request duration`,
		},
	},
}

func MustInit(address string, startPrometheusServer bool) {
	registrator, httpHandler, err := initPrometheus(packageName)
	if err != nil {
		log.Panic(err)
	}

	if startPrometheusServer {
		http.Handle("/metrics", httpHandler)
		go func() {
			//log.Debugf("prometheus server running on %s", address)
			log.Panic(http.ListenAndServe(address, nil))
		}()
	}

	for _, met := range counters {
		registrator.registerCounter(met)
	}

	for _, met := range gauges {
		registrator.registerGauge(met)
	}

	for _, met := range histograms {
		registrator.registerHistogram(met)
	}
}

func TrackTimeBegin() TimeTracker {
	return TimeTracker(time.Now())
}

func (t TimeTracker) Flush(ctx context.Context, metric uint8, attrs ...api.RecordOption) {
	RecordHist(ctx, metric, time.Since(time.Time(t)).Seconds(), attrs...)
}

func RecordHist(ctx context.Context, metric uint8, val float64, attrs ...api.RecordOption) {
	if m, ok := histograms[metric]; ok {
		m.fn(ctx, val, attrs...)
	} else {
		log.Panicf("meter %d is not histogram", metric)
	}
}

func UpdateGauge(ctx context.Context, metric uint8, val float64, attrs ...api.ObserveOption) {
	if m, ok := gauges[metric]; ok {
		m.fn(ctx, val, attrs...)
	} else {
		log.Panicf("meter %d is not gauge", metric)
	}
}

func IncCounter(ctx context.Context, metric uint8, attrs ...api.AddOption) {
	if m, ok := counters[metric]; ok {
		m.fn(ctx, 1, attrs...)
	} else {
		log.Panicf("meter %d is not counter", metric)
	}
}

func KeyValue(k attribute.Key, v attribute.Value) attribute.KeyValue {
	return attribute.KeyValue{Key: k, Value: v}
}
