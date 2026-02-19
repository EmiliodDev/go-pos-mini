package main

import (
	"net/http"
	"time"

	"github.com/EmiliodDev/go-pos/internal/healthcheck"
)

func (a *app) mount() http.Handler {
	mux := http.NewServeMux()

	healthcheck := healthcheck.NewHandler()
	mux.HandleFunc("/healthcheck", healthcheck.Healthcheck)

	return mux
}

func (a *app) run(h http.Handler) error {
	srv := &http.Server{
		Addr:         a.config.addr,
		Handler:      h,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	return srv.ListenAndServe()
}

type config struct {
	addr string
}

type app struct {
	config config
}
