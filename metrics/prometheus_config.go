package metrics

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	myapp = "apple_backend"
	// appVersion string
	// version    = prometheus.NewGauge(prometheus.GaugeOpts{
	// 	Name: "version",
	// 	Help: "Version information about this binary",
	// 	ConstLabels: map[string]string{
	// 		"version": appVersion,
	// 	},
	// })

	// Metric that count the processed operations
	opsProcessed = prometheus.NewCounter(prometheus.CounterOpts{
		Name: myapp + "_ops_total",
		Help: "The total number of processed events",
	})

	// Metric that count the total of request
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: myapp + "_http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "endpoint", "status"},
	)

	// Metric for measuring request duration
	httpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    myapp + "_http_request_duration_seconds",
			Help:    "Duration of HTTP requests in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "endpoint"},
	)

	// Metric to register errores
	errorCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: myapp + "_error_count_total",
			Help: "Total number of errors logged",
		},
		[]string{"type"},
	)
)

// Increment the metrics of the request HTTP
func RecordHTTPRequests(method, endpoint string, statusCode int, duration float64) {
	httpRequestsTotal.WithLabelValues(method, endpoint, http.StatusText(statusCode)).Inc()
	httpRequestDuration.WithLabelValues(method, endpoint).Observe(duration)
}

// Increment the metrics of errors
func RecordError(errorType string) {
	errorCount.WithLabelValues(errorType).Inc()
}

// Handler to get and view metrics
func GetMetricsHandler(r *prometheus.Registry) http.Handler {
	return promhttp.HandlerFor(r, promhttp.HandlerOpts{})
}

// Starts metric recording
func InitMetrics(r *prometheus.Registry) {
	log.Print("Registering Prometheus metrics...")

	r.MustRegister(
		opsProcessed,
		httpRequestsTotal,
		httpRequestDuration,
		errorCount,
	)
}
