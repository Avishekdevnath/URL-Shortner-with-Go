package redis

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"Backend/internal/models" // Local import for models
	"Backend/internal/store"  // Local import for store interface
)

// RedisStore represents a Redis-based store.
type RedisStore struct {
	client *redis.Client
}

// New creates a new RedisStore and connects to the Redis server.
func New(addr, password, db string) (*RedisStore, error) {
	// Initialize a Redis client
	client := redis.NewClient(&redis.Options{
		Addr:     addr,     // Redis server address (e.g., "localhost:6379")
		Password: password, // Redis password (default is no password)
		DB:       0,        // Default database
	})

	// Test the connection to Redis
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %v", err)
	}

	return &RedisStore{client: client}, nil
}

// Save stores the original URL and returns a shortened URL in Redis.
func (r *RedisStore) Save(url *models.ShortURLRequest) (*models.ShortURLResponse, error) {
	// Generate a unique short code (could be improved with a better generation strategy)
	shortCode := fmt.Sprintf("short%d", len(url.URL))

	// Store the original URL in Redis with the short code as the key
	err := r.client.Set(context.Background(), shortCode, url.URL, 0).Err()
	if err != nil {
		return nil, fmt.Errorf("failed to save URL in Redis: %v", err)
	}

	// Return the shortened URL response
	return &models.ShortURLResponse{ShortURL: shortCode}, nil
}

// Get retrieves the original URL for a given short code from Redis.
func (r *RedisStore) Get(shortCode string) (*models.ShortURLResponse, error) {
	// Get the original URL from Redis using the short code
	originalURL, err := r.client.Get(context.Background(), shortCode).Result()
	if err == redis.Nil {
		return nil, errors.New("short URL not found in Redis")
	}
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve URL from Redis: %v", err)
	}

	// Return the original URL in the response
	return &models.ShortURLResponse{ShortURL: originalURL}, nil
}

// Ensure RedisStore implements the Store interface
var _ store.Store = &RedisStore{}
