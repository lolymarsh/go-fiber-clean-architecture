package main

import (
	"log"
	"loly_api/config"
	"loly_api/server"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	server := server.NewServer(conf)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		if err := server.Run(); err != nil {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	<-quit
	log.Println("Received shutdown signal, initiating graceful shutdown...")

	if err := server.Shutdown(); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}

	log.Println("Server gracefully stopped")
}
