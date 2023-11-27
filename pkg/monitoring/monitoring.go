package monitoring

import (
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type PrometheusMetrics struct {
	RequestCounter prometheus.Counter
	Hits           *prometheus.CounterVec
	Duration       *prometheus.HistogramVec
}

func RegisterMonitoring(router *mux.Router) *PrometheusMetrics {
	var metrics = new(PrometheusMetrics)
	metrics.RequestCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "api_",
		Name:      "request_",
		Help:      "Number of request.",
	})
	metrics.Hits = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "api_",
		Name:      "hits_",
		Help:      "All request hits.",
	}, []string{"status", "path", "method"})
	metrics.Duration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "api_",
		Name:      "duration_",
		Help:      "Request duration.",
	}, []string{"status", "path", "method"})

	prometheus.MustRegister(metrics.Hits, metrics.Duration)

	router.Path("/metrics").Handler(promhttp.Handler())

	return metrics
}
