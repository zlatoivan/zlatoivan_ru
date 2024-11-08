package main

import (
	"context"
	"fmt"
	"log"
	"os/signal"
	"syscall"

	"github.com/zlatoivan/zlatoivan_ru/internal/config"
	"github.com/zlatoivan/zlatoivan_ru/internal/pkg/server"
)

func main() {
	ctx := context.Background()

	err := bootstrap(ctx)
	if err != nil {
		log.Fatalf("[main] bootstrap: %v", err)
	}
}

func bootstrap(ctx context.Context) error {
	ctx, cancel := signal.NotifyContext(ctx, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	cfg, err := config.New()
	if err != nil {
		return fmt.Errorf("config.New: %w", err)
	}

	httpServer := server.New()

	err = httpServer.Run(ctx, cfg.HTTPServer)
	if err != nil {
		return fmt.Errorf("httpServer.Run: %w", err)
	}

	<-ctx.Done()

	return nil
}
