package main

import (
	"fmt"
	"log"
	"loly_api/config"
	"loly_api/database"
	"loly_api/server"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	conf := config.LoadConfig()
	db := database.NewDatabase(conf)
	server := server.NewServer(conf, db)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		if err := server.Run(); err != nil {
			panic(fmt.Errorf("server failed to start: %w", err))
		}
	}()

	<-quit
	log.Println("Received shutdown signal, initiating graceful shutdown...")

	if err := server.Shutdown(); err != nil {
		panic(fmt.Errorf("server failed to shutdown: %w", err))
	}

	if err := db.Close(); err != nil {
		panic(fmt.Errorf("failed to close database connection: %w", err))
	}

	log.Println("Server gracefully stopped")
}
