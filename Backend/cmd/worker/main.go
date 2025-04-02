package main

import (
	"log"
	"time"
	"Backend/config" // Local import for the config package
	"Backend/internal/service" // Local import for the service layer
	"Backend/internal/store/postgres" // Local import for PostgreSQL store
)

func main() {
	// Load the configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Initialize PostgreSQL store
	postgresStore, err := postgres.New(cfg.PostgresDSN)
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}

	// Initialize the service with the PostgreSQL store
	urlService := service.NewURLService(postgresStore)

	// Run the worker
	runWorker(urlService)
}

// runWorker is where the background processing logic can be implemented
func runWorker(service *service.URLService) {
	// For example, the worker can run every 5 seconds
	for {
		// Example task: print a log to show that the worker is running
		log.Println("Worker is running...")

		// Add your background tasks here
		// For example, you might want to clean up expired URLs or perform analytics tasks

		// Sleep for a set interval (e.g., 5 seconds)
		time.Sleep(5 * time.Second)
	}
}
