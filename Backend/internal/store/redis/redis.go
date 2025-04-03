// S:\SDE\Hard Core\Learn\Golang\Projects\URL-Shortner-with-Go\Backend\internal\store\redis\redis.go
package redis

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"Backend/internal/models" // Local import for models
	"Backend/internal/store"  // Local import for store interface
	"math/rand"
	"time"
)

// RedisStore represents a Redis-based store.
type RedisStore struct {
	client *redis.Client
}

// GenerateRandomString generates a random string of a specified length
func GenerateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seed := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[seed.Intn(len(charset))]
	}
	return string(result)
}

// New creates a new RedisStore and connects to the Redis server.
func New(addr, password string, db int) (*RedisStore, error) {
	// Initialize a Redis client
	client := redis.NewClient(&redis.Options{
		Addr:     addr,     // Redis server address (e.g., "localhost:6379")
		Password: password, // Redis password (default is no password)
		DB:       db,       // Database number
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
	// Generate a random short code
	shortCode := GenerateRandomString(6)

	// Store the original URL in Redis with the short code as the key
	err := r.client.Set(context.Background(), shortCode, url.URL, 0).Err()
	if err != nil {
		return nil, fmt.Errorf("failed to save URL in Redis: %v", err)
	}

	// Return the shortened URL response with full URL
	return &models.ShortURLResponse{ShortURL: "http://localhost:8080/" + shortCode, OriginalURL: url.URL}, nil
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

	// Return the original URL in the response with full URL
	return &models.ShortURLResponse{ShortURL: "http://localhost:8080/" + shortCode, OriginalURL: originalURL}, nil
}

// GetAll retrieves all stored short URLs from Redis.
func (r *RedisStore) GetAll() (map[string]string, error) {
	ctx := context.Background()
	result := make(map[string]string)

	// Use SCAN to iterate over all keys in Redis
	var cursor uint64
	for {
		keys, nextCursor, err := r.client.Scan(ctx, cursor, "*", 10).Result()
		if err != nil {
			return nil, fmt.Errorf("failed to scan keys in Redis: %v", err)
		}

		// Fetch values for the retrieved keys
		for _, key := range keys {
			value, err := r.client.Get(ctx, key).Result()
			if err == redis.Nil {
				continue // Skip keys with no associated value
			}
			if err != nil {
				return nil, fmt.Errorf("failed to retrieve value for key %s: %v", key, err)
			}
			result[key] = value
		}

		// Break if the cursor is 0 (end of iteration)
		if nextCursor == 0 {
			break
		}
		cursor = nextCursor
	}

	return result, nil
}

// Ensure RedisStore implements the Store interface
var _ store.Store = &RedisStore{}