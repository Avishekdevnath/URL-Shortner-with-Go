// S:\SDE\Hard Core\Learn\Golang\Projects\URL-Shortner-with-Go\Backend\internal\store\memory\memory.go
package memory

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"Backend/internal/models" // Local import for models
)

// MemoryStore is an in-memory implementation of the Store interface.
type MemoryStore struct {
	mu    sync.RWMutex
	store map[string]string
}

// New creates a new in-memory store.
func New() *MemoryStore {
	return &MemoryStore{
		store: make(map[string]string),
	}
}

func GenerateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seed := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[seed.Intn(len(charset))]
	}
	fmt.Println("Random string", string(result))
	return string(result)
}

// Save stores the original URL and returns a shortened URL.
func (m *MemoryStore) Save(url *models.ShortURLRequest) (*models.ShortURLResponse, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Generate a completely random short code
	shortCode := GenerateRandomString(6) // Generate a 6-character random string
	fmt.Println("Short code", shortCode)
	m.store[shortCode] = url.URL

	return &models.ShortURLResponse{ShortURL: shortCode, OriginalURL: url.URL}, nil
}

// Get retrieves the original URL for the given short code.
func (m *MemoryStore) Get(shortCode string) (*models.ShortURLResponse, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	originalURL, exists := m.store[shortCode]
	if !exists {
		return nil, errors.New("short URL not found")
	}

	return &models.ShortURLResponse{ShortURL: "http://localhost:8080/" + shortCode, OriginalURL: originalURL}, nil
}

// GetAll retrieves all stored URLs as a map of short codes to original URLs.
func (m *MemoryStore) GetAll() (map[string]string, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	// Return a copy of the store to avoid external modification
	result := make(map[string]string)
	for shortCode, originalURL := range m.store {
		result[shortCode] = originalURL
	}
	return result, nil
}