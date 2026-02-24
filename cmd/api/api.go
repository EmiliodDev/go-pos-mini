package main

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/EmiliodDev/go-pos/internal/healthcheck"
	"github.com/jackc/pgx/v5"
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

	a.logger.Info("server listening on", "port", a.config.addr)
	return srv.ListenAndServe()
}

type app struct {
	config config
	db     *pgx.Conn
	logger *slog.Logger
}

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	dns string
}
