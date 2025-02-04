package main

import (
	"net/http"
	"os"

	"time"

	"github.com/alex-arraga/apple_store/metrics"
	"github.com/alex-arraga/apple_store/middlewares"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = zerolog.New(os.Stderr).With().Timestamp().Logger()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal().Msg("PORT not set in environment variables")
	}

	// Prometheus starts to record metrics
	rec := prometheus.NewRegistry()
	metrics.InitMetrics(rec)

	router := chi.NewRouter()
	router.Use(middlewares.RecordPrometheusMiddleware)

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Hellow world"))
		if err != nil {
			log.Fatal().Msgf("Error in '/' route: %s", err)
		}
	})

	router.Get("/metrics", func(w http.ResponseWriter, r *http.Request) {
		promhttp.HandlerFor(rec, promhttp.HandlerOpts{})
	})

	s := &http.Server{
		Handler:      router,
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Print("Server listening")
	err := s.ListenAndServe()
	if err != nil {
		log.Fatal().Msgf("Server failed: %s", err)
	}
}
