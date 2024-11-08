package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/zlatoivan/zlatoivan_ru/internal/config"
)

const colon = ":"

// Server - интерфейс http-сервера
type Server interface {
	Run(ctx context.Context, cfg config.HTTPServer) error
}

type server struct{}

// New - создает http-сервер
func New() Server {
	return server{}
}

// Run - метод запуска http-сервера
func (s server) Run(ctx context.Context, cfg config.HTTPServer) error {
	router := createRouter()
	httpServer := http.Server{
		Addr:              colon + cfg.Port,
		Handler:           router,
		ReadHeaderTimeout: cfg.ReadHeaderTimeout,
		ReadTimeout:       cfg.ReadTimeout,
		WriteTimeout:      cfg.WriteTimeout,
		IdleTimeout:       cfg.IdleTimeout,
	}

	go httpServerRun(&httpServer)

	go gracefulShutdown(ctx, &httpServer)

	log.Printf("[httpServer] starting on %s\n", cfg.Port)

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
