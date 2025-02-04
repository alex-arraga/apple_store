package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	log.Print("Getting PORT")
	const envFile string = ".env"

	// Load enviroment variables from the .env file
	if err := godotenv.Load(envFile); err != nil {
		log.Fatalf("failed to load .env file: %v", err)
	}

	port := os.Getenv("PORT")
	log.Printf("PORT: %v", port)

	router := chi.NewRouter()

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
			log.Fatalf("Error in '/' route: %s", err)
		}
	})

	s := &http.Server{
		Handler:      router,
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	err := s.ListenAndServe()
	if err != nil {
		log.Fatalf("Server failed: %s", err)
	}
	log.Print("Server listening")
}
