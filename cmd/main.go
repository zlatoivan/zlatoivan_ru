package main

import (
	"context"
	"fmt"
	"log"
	"os/signal"
	"syscall"

	"zlatoivan_ru/internal/config"
	"zlatoivan_ru/internal/server"
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

	mainServer := server.New()

	mainServer.Run(ctx, cfg.Server)
	if err != nil {
		return fmt.Errorf("mainServer.Run: %w", err)
	}

	<-ctx.Done()

	return nil
}
