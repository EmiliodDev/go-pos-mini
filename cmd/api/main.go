package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx := context.Background()

	dbConfig := dbConfig()

	cfg := config{
		addr: ":8080",
		db:   *dbConfig,
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	connPool, err := pgxpool.NewWithConfig(ctx, &cfg.db)
	if err != nil {
		logger.Error("error while creating connection to database")
	}

	logger.Info("connected to database")

	api := app{
		config: cfg,
		db:     connPool,
		logger: logger,
	}

	if err := api.run(api.mount()); err != nil {
		slog.Error("server failed to start", "error", err)
		os.Exit(1)
	}
}
