package router

import (
	"github.com/gin-gonic/gin"
	"Backend/internal/api/handlers"  // Local import for handlers
	"Backend/internal/api/middleware" // Local import for middleware
	"Backend/internal/service"       // Local import for service layer
)

// SetupRouter sets up the Gin routes and handlers. Now accepts urlService.
func SetupRouter(urlService *service.URLService) *gin.Engine {
	router := gin.Default()

	// Use the logging middleware
	router.Use(middleware.Logger())

	// Define the routes and pass urlService to handlers
	router.POST("/shorten", func(c *gin.Context) {
		handlers.ShortenURL(c, urlService)    // Pass urlService to ShortenURL
	})
	router.GET("/short/:code", func(c *gin.Context) {
		handlers.RedirectURL(c, urlService) // Pass urlService to RedirectURL
	})

	return router
}
