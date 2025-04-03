// S:\SDE\Hard Core\Learn\Golang\Projects\URL-Shortner-with-Go\Backend\pkg\logger\logger.go
package logger

import (
	"log"
)

// Info logs an informational message.
func Info(msg string) {
	log.Printf("[INFO] %s", msg)
}

// Error logs an error message.
func Error(msg string) {
	log.Printf("[ERROR] %s", msg)
}