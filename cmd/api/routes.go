package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) routes() http.Handler {
	r := chi.NewRouter()
	r.NotFound(http.HandlerFunc(app.notFoundResponse))
	r.MethodNotAllowed(http.HandlerFunc(app.methodNotAllowedResponse))
	r.Route("/v1", func(r chi.Router) {
		r.Get("/healthcheck", app.healthcheckHandler)

		r.Route("/movies", func(r chi.Router) {
			r.Post("/", app.createMovieHandler)
			r.Get("/{id}", app.showMovieHandler)
		})
	})

	return r
}
