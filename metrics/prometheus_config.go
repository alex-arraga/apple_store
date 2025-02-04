package metrics

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
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

	// Métrica para contar operaciones procesadas
	opsProcessed = prometheus.NewCounter(prometheus.CounterOpts{
		Name: myapp + "_ops_total",
		Help: "The total number of processed events",
	})

	// Métrica para la cantidad de solicitudes HTTP
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: myapp + "_http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "endpoint", "status"},
	)

	// Métrica para medir la duración de las solicitudes HTTP
	httpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    myapp + "_http_request_duration_seconds",
			Help:    "Duration of HTTP requests in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "endpoint"},
	)

	// Métrica para registrar errores
	errorCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: myapp + "_error_count_total",
			Help: "Total number of errors logged",
		},
		[]string{"type"},
	)
)

// Incrementa las métricas de solicitudes HTTP
func RecordHTTPRequests(method, endpoint string, statusCode int, duration float64) {
	httpRequestsTotal.WithLabelValues(method, endpoint, http.StatusText(statusCode)).Inc()
	httpRequestDuration.WithLabelValues(method, endpoint).Observe(duration)
}

// Incrementa las métricas de errores
func RecordError(errorType string) {
	errorCount.WithLabelValues(errorType).Inc()
}

func InitMetrics(r *prometheus.Registry) {
	// Record metrics
	log.Print("Starting Prometheus...")

	r.MustRegister(opsProcessed)
	r.MustRegister(httpRequestsTotal)
	r.MustRegister(httpRequestDuration)
	r.MustRegister(errorCount)
}
