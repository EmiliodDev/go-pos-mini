package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/EmiliodDev/go-pos/cmd/internal/healthcheck"
)

const (
	_shutdownPeriod      = 15 * time.Second
	_shutdownHardPeriod  = 3 * time.Second
	_readinessDrainDelay = 5 * time.Second
)

var isShutingDown atomic.Bool

func (a *app) mount() http.Handler {
	mux := http.NewServeMux()

	healthcheck := healthcheck.NewHandler()
	mux.HandleFunc("/healthcheck", healthcheck.Healthcheck)

	return mux
}

func (a *app) run(h http.Handler) {
	rootCtx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	ongoingCtx, stopOngoingCtx := context.WithCancel(context.Background())
	srv := &http.Server{
		Addr:         a.config.addr,
		Handler:      h,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
		BaseContext: func(l net.Listener) context.Context {
			return ongoingCtx
		},
	}

	go func() {
		log.Printf("Server starting on %s", a.config.addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	<-rootCtx.Done()
	stop()
	isShutingDown.Store(true)
	log.Println("Received shutdown signal, shutting down.")

	time.Sleep(_readinessDrainDelay)
	log.Println("Readiness check propagated, now waiting for ongoing request to finish.")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), _shutdownPeriod)
	defer cancel()
	err := srv.Shutdown(shutdownCtx)
	stopOngoingCtx()
	if err != nil {
		log.Println("Failed to wait for ongoing request to finish, waiting for forced cancellation.")
		time.Sleep(_shutdownHardPeriod)
	}

	log.Println("Server shutdown gracefully.")
}

type config struct {
	addr string
}

type app struct {
	config config
}
