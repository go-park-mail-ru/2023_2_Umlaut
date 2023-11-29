package monitoring

import (
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type PrometheusMetrics struct {
	Hits     *prometheus.CounterVec
	Duration *prometheus.HistogramVec
}

func RegisterMonitoring(router *mux.Router) *PrometheusMetrics {
	var metrics = new(PrometheusMetrics)
	metrics.Hits = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "api",
		Name:      "hits_total",
		Help:      "Total number of hits.",
	}, []string{"path", "method"})
	metrics.Duration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "api",
		Name:      "request_duration_seconds",
		Help:      "Duration of requests.",
		Buckets:   prometheus.DefBuckets,
	}, []string{"status", "path", "method"})

	prometheus.MustRegister(metrics.Hits, metrics.Duration)

	router.Path("/metrics").Handler(promhttp.Handler())

	return metrics
}
