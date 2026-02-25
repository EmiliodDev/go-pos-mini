package main

import (
	"context"
	"fmt"
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
		databaseUrl              = "postgres://admin:admin@localhost:5432/go_pos"
	)

	dbConfig, err := pgxpool.ParseConfig(databaseUrl)
	if err != nil {
		log.Fatal("Failed to create a config, error: ", err)
	}

	dbConfig.MaxConns = defaultMaxConns
	dbConfig.MinConns = defaultMinConns
	dbConfig.MaxConnLifetime = defaultMaxConnLifetime
	dbConfig.MaxConnIdleTime = defaultMaxConnIdleTime
	dbConfig.HealthCheckPeriod = defaultHealthCheckPeriod
	dbConfig.ConnConfig.ConnectTimeout = defaultConnectTimeout

	dbConfig.BeforeConnect = func(ctx context.Context, c *pgx.ConnConfig) error {
		fmt.Println("before")
		return nil
	}

	dbConfig.AfterRelease = func(c *pgx.Conn) bool {
		log.Println("after release")
		return true
	}

	dbConfig.BeforeClose = func(c *pgx.Conn) {
		fmt.Println("before close")
	}

	return dbConfig
}
