package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com-test/internal/di"
)

const defaultPort = "8000"

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()

	server, err := di.Init(ctx)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		<-ctx.Done()

		server.Shutdown(ctx)
	}()

	port := os.Getenv("PORT")
	if strings.TrimSpace(port) == "" {
		port = defaultPort
	}

	log.Fatal(server.Start(":" + port))
}
