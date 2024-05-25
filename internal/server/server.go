package server

import (
	"context"
	"log"
	"net/http"
	"sync"
	"time"

	"zlatoivan_ru/internal/config"
)

type balancerI interface {
	balance()
}

type Server struct {
	balancer balancerI
}

func New() Server {
	return Server{}
}

func (s Server) Run(ctx context.Context, cfg config.Server) {
	router := s.createRouter()
	httpServer := &http.Server{
		Addr:    "localhost:" + cfg.HttpPort,
		Handler: router,
	}

	wg := sync.WaitGroup{}

	log.Printf("[httpServer] starting on %s\n", cfg.HttpPort)

	wg.Add(1)
	go func() {
		httpServerRun(httpServer)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		gracefulShutdown(ctx, httpServer)
		wg.Done()
	}()

	wg.Wait()
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
