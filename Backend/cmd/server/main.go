package main

import (
	"log"
	"github.com/Avishekdevnath/URL-Shortner-with-Go/Backend/internal/api/router"
)

func main() {
	// Initialize the router
	r := router.SetupRouter()

	// Start the server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
