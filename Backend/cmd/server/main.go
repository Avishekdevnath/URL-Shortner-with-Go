package main

import (
	"log"
	"Backend/config"         // Local import for config
	"Backend/internal/api/router"  // Local import for router
	"Backend/internal/service"    // Local import for service layer
	"Backend/internal/store/memory" // Local import for memory store
)

func main() {
	// Load the configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Initialize memory store (or replace with other stores like PostgreSQL or Redis)
	postgresStore, err := memory.New() // Use `postgres.New(cfg.PostgresDSN)` for PostgreSQL
	if err != nil {
		log.Fatalf("Failed to connect to memory store: %v", err)
	}

	// Initialize the service with the store (memory store or any other store)
	urlService := service.NewURLService(postgresStore)

	// Setup the router with the URL service
	r := router.SetupRouter(urlService)  // Passing urlService to router

	// Start the server on the configured port
	if err := r.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
