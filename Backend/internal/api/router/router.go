// S:\SDE\Hard Core\Learn\Golang\Projects\URL-Shortner-with-Go\Backend\internal\api\router\router.go
package router

import (
	"github.com/gin-gonic/gin"
	"Backend/internal/api/handlers"  // Local import for handlers
	"Backend/internal/api/middleware" // Local import for middleware
	"Backend/internal/service"       // Local import for service layer
)

// SetupRouter sets up the Gin routes and handlers. Now includes /urls to view all URLs.
func SetupRouter(router *gin.Engine, urlService *service.URLService) {
	// Use the logging middleware
	router.Use(middleware.Logger())
	// Use the CORS middleware
	router.Use(middleware.CORS())

	// Define the routes and pass urlService to handlers
	router.POST("/api/shorten", func(c *gin.Context) {
		handlers.ShortenURL(c, urlService)    // Pass urlService to ShortenURL
	})
	router.GET("/:code", func(c *gin.Context) {
		handlers.RedirectURL(c, urlService) // Pass urlService to RedirectURL
	})
	router.GET("/urls", func(c *gin.Context) {
		handlers.GetAllURLs(c, urlService)  // Pass urlService to GetAllURLs
	})
}