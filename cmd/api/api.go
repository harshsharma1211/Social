package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type application struct {
	config config
}

type config struct {
	addr string
}

func (app *application) mount() http.Handler {

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.HealthCheckHandler)
	})

	return r

}

func (app *application) run(handler http.Handler) error {

	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      handler,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 20,
		IdleTimeout:  time.Minute,
	}

	log.Printf("Starting server at %s \n", app.config.addr)

	return srv.ListenAndServe()

}
