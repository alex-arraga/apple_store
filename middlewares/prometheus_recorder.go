package middlewares

import (
	"net/http"
	"time"

	"github.com/alex-arraga/apple_store/metrics"
)

// PrometheusMiddleware mide y registra las métricas de las solicitudes HTTP
func RecordPrometheusMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Envolver el ResponseWriter para capturar el código de estado
		rec := &responseRecorder{ResponseWriter: w, statusCode: http.StatusOK}
		next.ServeHTTP(rec, r)

		duration := time.Since(start).Seconds()
		metrics.RecordHTTPRequests(r.Method, r.URL.Path, rec.statusCode, duration)
	})
}

// responseRecorder es un wrapper para capturar el código de estado
type responseRecorder struct {
	http.ResponseWriter
	statusCode int
}
