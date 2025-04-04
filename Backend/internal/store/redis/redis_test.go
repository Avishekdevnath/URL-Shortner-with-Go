package redis

import (
	"context"
	"testing"
	"Backend/internal/models"
	"github.com/go-redis/redis/v8"
)

func setupRedis(t *testing.T) *RedisStore {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		t.Fatalf("Failed to connect to Redis: %v", err)
	}
	return &RedisStore{client: client}
}

func TestRedisStore(t *testing.T) {
	store := setupRedis(t)

	// Test Save
	req := &models.ShortURLRequest{URL: "https://example.com"}
	resp, err := store.Save(req)
	if err != nil {
		t.Fatalf("Save failed: %v", err)
	}
	if resp.ShortURL == "" || resp.OriginalURL != req.URL {
		t.Errorf("Save returned invalid response: %+v", resp)
	}

	// Test Get
	got, err := store.Get(resp.ShortURL)
	if err != nil {
		t.Fatalf("Get failed: %v", err)
	}
	if got.OriginalURL != req.URL {
		t.Errorf("Get returned wrong URL: got %s, want %s", got.OriginalURL, req.URL)
	}

	// Test GetAll
	urls, err := store.GetAll()
	if err != nil {
		t.Fatalf("GetAll failed: %v", err)
	}
	if len(urls) != 1 || urls[0].OriginalURL != req.URL {
		t.Errorf("GetAll returned wrong data: %+v", urls)
	}
	if urls[0].UserID != nil {
		t.Errorf("UserID should be nil for anonymous: got %v", urls[0].UserID)
	}
}