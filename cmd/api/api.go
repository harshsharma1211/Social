package main

import (
	"log"
	"net/http"
	"time"
)

type application struct {
	config config
}

type config struct {
	addr string
}

func (app *application) mount() *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("GET v1/health", app.HealthCheckHandler)

	return mux

}

func (app *application) run(mux *http.ServeMux) error {

	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 20,
		IdleTimeout:  time.Minute,
	}

	log.Printf("Starting server at %s \n", app.config.addr)

	return srv.ListenAndServe()

}
