package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/zlatoivan/zlatoivan_ru/internal/config"
)

const colon = ":"

type Server interface {
	Run(ctx context.Context, cfg config.Server) error
}

type server struct{}

func New() Server {
	return server{}
}

func (s server) Run(ctx context.Context, cfg config.Server) error {
	router := createRouter()
	httpServer := http.Server{
		Addr:              colon + cfg.HTTPPort,
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	log.Printf("[httpServer] starting on %s\n", cfg.HTTPPort)

	go func() {
		httpServerRun(&httpServer)
	}()

	go func() {
		gracefulShutdown(ctx, &httpServer)
	}()

	return nil
}

func httpServerRun(httpServer *http.Server) {
	err := httpServer.ListenAndServe()
	if err != nil {
		log.Printf("[httpServer] ListenAndServe: %v\n", err)
	}
}

func gracefulShutdown(ctx context.Context, httpServer *http.Server) {
	<-ctx.Done()
	ctxTo, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("[gracefulShutdown] shutting down")

	err := httpServer.Shutdown(ctxTo)
	if err != nil {
		log.Printf("httpServer.Shutdown: %v\n", err)
	}

	log.Println("[gracefulShutdown] shut down successfully")
}
