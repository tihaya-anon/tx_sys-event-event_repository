package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/robfig/cron/v3"
	constant_redis "github.com/tihaya-anon/tx_sys-event-event_repository/src/constant/redis"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/listener"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/server"
)

func main() {
	// Initialize server
	srv, err := server.NewServer(50051)
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}
	go srv.Start()

	// Initialize Redis client
	rdb := redis.NewClient(&redis.Options{
		Addr: constant_redis.REDIS_ADDR,
	})
	if _, err := rdb.Ping(context.Background()).Result(); err != nil {
		log.Printf("Warning: Redis connection failed: %v", err)
	}

	// Initialize consumer manager
	ctx := context.Background()
	listener.InitConsumerManager(ctx, rdb)

	// Set up cron job for listener
	c := cron.New()
	c.AddFunc("0 */1 * * * *", func() {
		listener.CreateListener(context.Background(), nil, rdb)
	})
	go c.Start()
	// Wait for interrupt signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	<-sigChan
	log.Println("Shutting down services...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown services in order
	c.Stop()
	listener.ShutdownConsumerManager(ctx)
	srv.Shutdown(ctx)

	// Close Redis connection
	if err := rdb.Close(); err != nil {
		log.Printf("Error closing Redis connection: %v", err)
	}
}
