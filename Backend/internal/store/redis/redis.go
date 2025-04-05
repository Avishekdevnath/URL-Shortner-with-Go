// S:\SDE\Hard Core\Learn\Golang\Projects\URL-Shortner-with-Go\Backend\internal\store\redis\redis.go

package redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"time"
	"Backend/internal/models"
	"Backend/config"
	"github.com/go-redis/redis/v8"
	
)

type RedisStore struct {
	client *redis.Client
}

func GenerateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seed := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[seed.Intn(len(charset))]
	}
	return string(result)
}

// New initializes the Redis store with the URL from config.
func New() (*RedisStore, error) {
	cfg, err := config.LoadConfig()  // Load the configuration
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %v", err)
	}

	client := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisURL,  // Use Redis URL from config
		Password: "",             // Use empty password if not needed
		DB:       0,              // Use default DB (0)
	})

	_, err = client.Ping(context.Background()).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %v", err)
	}
	return &RedisStore{client: client}, nil
}

// Save stores the URL in Redis cache.
func (r *RedisStore) Save(url *models.ShortURLRequest) (*models.ShortURLResponse, error) {
	shortCode := GenerateRandomString(6)
	createdAt := time.Now().Unix()
	newURL := models.URL{
		ShortCode:   shortCode,
		OriginalURL: url.URL,
		CreatedAt:   createdAt,
	}

	data, err := json.Marshal(newURL)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal URL: %v", err)
	}
	ctx := context.Background()
	// Store the shortened URL in Redis
	if err := r.client.Set(ctx, shortCode, data, 0).Err(); err != nil {
		return nil, fmt.Errorf("failed to save URL in Redis: %v", err)
	}
	// Add to ZSET for ordered retrieval based on creation time
	if err := r.client.ZAdd(ctx, "urls", &redis.Z{Score: float64(createdAt), Member: shortCode}).Err(); err != nil {
		return nil, fmt.Errorf("failed to add to ZSET: %v", err)
	}

	return &models.ShortURLResponse{ShortURL: shortCode, OriginalURL: url.URL}, nil
}

// Get retrieves the URL from Redis.
func (r *RedisStore) Get(shortCode string) (*models.ShortURLResponse, error) {
	ctx := context.Background()
	data, err := r.client.Get(ctx, shortCode).Result()
	if err == redis.Nil {
		return nil, errors.New("short URL not found in Redis")
	}
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve URL from Redis: %v", err)
	}

	var url models.URL
	if err := json.Unmarshal([]byte(data), &url); err != nil {
		return nil, fmt.Errorf("failed to unmarshal URL: %v", err)
	}

	return &models.ShortURLResponse{ShortURL: shortCode, OriginalURL: url.OriginalURL}, nil
}

// GetAll retrieves all URLs from Redis.
func (r *RedisStore) GetAll() ([]models.URL, error) {
	ctx := context.Background()
	shortCodes, err := r.client.ZRevRange(ctx, "urls", 0, -1).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve URLs from ZSET: %v", err)
	}

	var urls []models.URL
	for _, shortCode := range shortCodes {
		data, err := r.client.Get(ctx, shortCode).Result()
		if err == redis.Nil {
			continue
		}
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve URL %s: %v", shortCode, err)
		}

		var url models.URL
		if err := json.Unmarshal([]byte(data), &url); err != nil {
			return nil, fmt.Errorf("failed to unmarshal URL %s: %v", shortCode, err)
		}
		urls = append(urls, url)
	}

	return urls, nil
}

// Health checks if Redis is responsive.
func (r *RedisStore) Health() error {
	_, err := r.client.Ping(context.Background()).Result()
	return err
}
