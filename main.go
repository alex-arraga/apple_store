package main

import (
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/alex-arraga/apple_store/db/conn"
	"github.com/alex-arraga/apple_store/metrics"
	"github.com/alex-arraga/apple_store/middlewares"
)

type User struct {
	Name  string
	Email string
}

func main() {
	// ! Config zerolog to promtail
	logFile, err := os.OpenFile("/app/logs/app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to open log file")
	}
	defer logFile.Close()

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: logFile})
	log.Print("Starting log app...")

	// Get env
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal().Msg("PORT not set in environment variables")
	}

	db, err := conn.InitDB()
	if err != nil {
		log.Fatal().Msgf("Failed DB connection: %v", err)
	}

	// Auto-migrate to create table if it doesn't exist
	if err := db.AutoMigrate(&User{}); err != nil {
		log.Fatal().Msgf("Error in migration: %v", err)
	}

	// Insert a test user
	user := User{Name: "Juan Perez", Email: "juan.perez@example.com"}
	if err := db.Create(&user).Error; err != nil {
		log.Fatal().Msgf("Failed to create user: %v", err)
	}

	log.Printf("User successfully created: %v \n", user)

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

	router.Handle("/metrics", metrics.GetMetricsHandler(rec))

	s := &http.Server{
		Handler:      router,
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Print("Server listening...")
	err = s.ListenAndServe()
	if err != nil {
		log.Fatal().Msgf("Server failed: %s", err)
	}
}
