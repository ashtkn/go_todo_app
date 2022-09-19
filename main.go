package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/ashtkn/go_todo_app/config"
)

func run(ctx context.Context) error {
	// Load config
	cfg, err := config.New()
	if err != nil {
		return err
	}

	// Create listener
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Fatalf("failed to listen port %d: %v", cfg.Port, err)
	}

	url := fmt.Sprintf("http://%s", l.Addr().String())
	log.Printf("server is running at %s", url)

	mux := NewMux()
	s := NewServer(l, mux)

	return s.Run(ctx)
}

func main() {
	if err := run(context.Background()); err != nil {
		fmt.Printf("failed to terminate server: %v", err)
		os.Exit(1)
	}
}
