package middleware

import (
	"log"
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

		// Log the details of the request
		log.Printf("%s %s %d %v", method, path, statusCode, duration)
	}
}
