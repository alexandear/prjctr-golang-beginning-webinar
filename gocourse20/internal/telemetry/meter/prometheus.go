package meter

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/prometheus"
	api "go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/sdk/metric"
)

var UserKey = attribute.Key("user")

type prometheusRegistrator struct {
	meter api.Meter
}

func initPrometheus(name string) (*prometheusRegistrator, http.Handler, error) {
	exporter, err := prometheus.New()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to initialize prometheus exporter: %w", err)
	}

	provider := metric.NewMeterProvider(metric.WithReader(exporter))
	otel.SetMeterProvider(provider)

	return &prometheusRegistrator{provider.Meter(name)}, promhttp.Handler(), nil
}

func (r *prometheusRegistrator) registerHistogram(mc *histMetric) {
	hist, err := r.meter.Float64Histogram(mc.name, api.WithDescription(mc.description))
	if err != nil {
		log.Panicf("failed to initialize instrument: %v", err)
	}

	mc.fn = func(ctx context.Context, val float64, attrs ...api.RecordOption) {
		hist.Record(ctx, val, attrs...)
	}
}

func (r *prometheusRegistrator) registerCounter(mc *countMetric) {
	counter, err := r.meter.Float64Counter(mc.name, api.WithDescription(mc.description))
	if err != nil {
		log.Panicf("failed to initialize instrument: %v", err)
	}

	mc.fn = func(ctx context.Context, val float64, attrs ...api.AddOption) {
		counter.Add(ctx, val, attrs...)
	}
}

func (r *prometheusRegistrator) registerGauge(mc *gaugeMetric) {
	observerLock := new(sync.RWMutex)
	observerValueToReport := new(float64)
	observerAttrsToReport := new([]api.ObserveOption)

	gaugeObserver, err := r.meter.Float64ObservableGauge(mc.name, api.WithDescription(mc.description))
	if err != nil {
		log.Panic("failed to initialize instrument: %v", err)
	}

	_, _ = r.meter.RegisterCallback(func(ctx context.Context, o api.Observer) error {
		(*observerLock).RLock()
		value := *observerValueToReport
		attrs := *observerAttrsToReport
		(*observerLock).RUnlock()
		o.ObserveFloat64(gaugeObserver, value, attrs...)

		return nil
	}, gaugeObserver)

	mc.fn = func(ctx context.Context, val float64, attrs ...api.ObserveOption) {
		(*observerLock).Lock()
		*observerValueToReport += val
		*observerAttrsToReport = attrs
		(*observerLock).Unlock()
	}
}
