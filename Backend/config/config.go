// S:\SDE\Hard Core\Learn\Golang\Projects\URL-Shortner-with-Go\Backend\config\config.go
package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config holds the application's configuration settings.
type Config struct {
	// Database configuration
	PostgresDSN string
	// Redis configuration
	RedisURL     string // Redis URL
	RedisPassword string
	RedisDB       int
	// Server configuration
	ServerPort    string
	// Store type configuration (memory, postgres, redis)
	StoreType     string
	// Memory store configuration
	MemoryFilePath string // File path for MemoryStore persistence
}

// LoadConfig loads the environment variables from a .env file and returns the configuration.
func LoadConfig() (*Config, error) {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using default values")
	}

	// Return the loaded configuration
	return &Config{
		PostgresDSN:   getEnv("POSTGRES_DSN", "postgresql://localhost:5432/postgres"),
		RedisURL:      getEnv("REDIS_URL", "rediss://localhost:6379"),
		RedisPassword: getEnv("REDIS_PASSWORD", ""),
		RedisDB:       getEnvAsInt("REDIS_DB", 0),
		ServerPort:    getEnv("SERVER_PORT", "8080"),
		StoreType:     getEnv("STORE_TYPE", "memory"),
		MemoryFilePath: getEnv("MEMORY_FILE_PATH", "data.json"), // Default file path for MemoryStore
	}, nil
}

// getEnv gets an environment variable or returns a fallback value if not set
func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	return value
}

// getEnvAsInt gets an environment variable as an integer or returns a fallback value if not set
func getEnvAsInt(key string, fallback int) int {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	// Convert string to int
	intValue, err := strconv.Atoi(value)
	if err != nil {
		log.Printf("Error converting %s to int, using fallback: %v", key, err)
		return fallback
	}
	return intValue
}
