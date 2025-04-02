package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds the application's configuration settings.
type Config struct {
	// Database configuration
	PostgresDSN string
	// Redis configuration
	RedisAddr string
	// Server configuration
	ServerPort string
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
		PostgresDSN: getEnv("POSTGRES_DSN", "host=localhost user=postgres dbname=url_shortner password=1234 port=5432 sslmode=disable"),
		RedisAddr:   getEnv("REDIS_ADDR", "localhost:6379"),
		ServerPort:  getEnv("SERVER_PORT", "8080"),
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
