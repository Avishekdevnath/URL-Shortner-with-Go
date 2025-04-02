package router

import (
	"github.com/gin-gonic/gin"
	"github.com/Avishekdevnath/URL-Shortner-with-Go/Backend/internal/api/handlers"
)

// SetupRouter sets up the Gin routes and handlers.
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Define the routes
	router.POST("/shorten", handlers.ShortenURL)
	router.GET("/short/:code", handlers.RedirectURL)

	return router
}
