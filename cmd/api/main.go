package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
}

type application struct {
	config
	logger *log.Logger
}

func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 8080, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &application{
		config: cfg,
		logger: logger,
	}

	r := chi.NewRouter()
	r.Route("/v1", func(r chi.Router) {
		r.Get("/healthcheck", app.healthcheckHandler)
	})

	srv := &http.Server{
		Addr:        fmt.Sprintf(":%d", cfg.port),
		Handler:     r,
		IdleTimeout: time.Minute,
		ReadTimeout: 10 * time.Second,
	}

	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		logger.Fatal(err)
	}
}
