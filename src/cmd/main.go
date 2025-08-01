package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/tihaya-anon/tx_sys-event-event_repository/src/server"
)

func main() {
	srv, err := server.NewServer(50051)
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}
	go srv.Start()

	// Wait for interrupt signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	<-sigChan
	log.Println("Shutting down gRPC server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	srv.Shutdown(ctx)
}
