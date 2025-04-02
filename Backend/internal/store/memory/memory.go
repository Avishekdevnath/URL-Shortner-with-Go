package memory

import (
	"errors"
	"sync"

	"github.com/Avishekdevnath/URL-Shortner-with-Go/Backend/internal/models"
	"github.com/Avishekdevnath/URL-Shortner-with-Go/Backend/internal/store"
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

// Save stores the original URL and returns a shortened URL.
func (m *MemoryStore) Save(url *models.ShortURLRequest) (*models.ShortURLResponse, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	// For simplicity, generating a simple hash-like "short code" (you can improve this)
	shortCode := "short" + string(len(m.store)+1)
	m.store[shortCode] = url.URL

	return &models.ShortURLResponse{ShortURL: shortCode}, nil
}

// Get retrieves the original URL for the given short code.
func (m *MemoryStore) Get(shortCode string) (*models.ShortURLResponse, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	originalURL, exists := m.store[shortCode]
	if !exists {
		return nil, errors.New("short URL not found")
	}

	return &models.ShortURLResponse{ShortURL: originalURL}, nil
}
