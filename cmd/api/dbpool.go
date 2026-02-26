package main

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func dbConfig() *pgxpool.Config {
	const (
		defaultMaxConns          = int32(4)
		defaultMinConns          = int32(0)
		defaultMaxConnLifetime   = time.Hour
		defaultMaxConnIdleTime   = time.Minute * 15
		defaultHealthCheckPeriod = time.Minute
		defaultConnectTimeout    = time.Second * 5
		databaseURL              = "postgres://admin:admin@localhost:5432/go_pos"
	)

	dbConfig, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		log.Fatal("Failed to create a config, error: ", err)
	}

	dbConfig.MaxConns = defaultMaxConns
	dbConfig.MinConns = defaultMinConns
	dbConfig.MaxConnLifetime = defaultMaxConnLifetime
	dbConfig.MaxConnIdleTime = defaultMaxConnIdleTime
	dbConfig.HealthCheckPeriod = defaultHealthCheckPeriod
	dbConfig.ConnConfig.ConnectTimeout = defaultConnectTimeout

	// TODO: Investigate real implementations for before/after & before-close
	dbConfig.BeforeConnect = func(ctx context.Context, c *pgx.ConnConfig) error {
		log.Println("before")
		return nil
	}

	dbConfig.AfterRelease = func(c *pgx.Conn) bool {
		log.Println("after release")
		return true
	}

	dbConfig.BeforeClose = func(c *pgx.Conn) {
		log.Println("before close")
	}

	return dbConfig
}
