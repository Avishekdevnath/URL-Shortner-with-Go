// S:\SDE\Hard Core\Learn\Golang\Projects\URL-Shortner-with-Go\Backend\internal\api\middleware\middleware.go




package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Logger is a middleware that logs the details of each HTTP request.
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the current time before the request
		start := time.Now()

		// Process the request
		c.Next()

		// Log the request details after it is processed
		duration := time.Since(start)
		statusCode := c.Writer.Status()
		method := c.Request.Method
		path := c.Request.URL.Path
		clientIP := c.ClientIP()

		// Log the details of the request
		log.Printf("%s %s %d %v %s", method, path, statusCode, duration, clientIP)
	}
}

// CORS is a middleware that adds CORS headers to responses.
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*") // Allow all origins for testing
		c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Accept, Content-Type")

		// Handle preflight OPTIONS requests
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		// Process the request
		c.Next()
	}
}