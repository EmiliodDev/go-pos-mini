package main

import (
	"log/slog"
	"net/http"
	"time"

	repo "github.com/EmiliodDev/go-pos/internal/adapters/pgdb/sqlc"
	"github.com/EmiliodDev/go-pos/internal/healthcheck"
	"github.com/EmiliodDev/go-pos/internal/products"
	"github.com/jackc/pgx/v5/pgxpool"
)

func (a *app) mount() http.Handler {
	mux := http.NewServeMux()

	healthcheck := healthcheck.NewHandler()
	mux.HandleFunc("/healthcheck", healthcheck.Healthcheck)

	// repositories
	repo := repo.New(a.db)

	productsService := products.NewService(repo)
	productsHandler := products.NewHandler(productsService)

	// products endpoints
	mux.HandleFunc("GET /products", productsHandler.ListProducts)
	mux.HandleFunc("POST /product", productsHandler.CreateProduct)

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
	db     *pgxpool.Pool
	logger *slog.Logger
}

type config struct {
	addr string
	db   pgxpool.Config
}
