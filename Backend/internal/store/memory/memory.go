// S:\SDE\Hard Core\Learn\Golang\Projects\URL-Shortner-with-Go\Backend\internal\store\memory\memory.go
package memory

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
	"Backend/internal/models"
)

type MemoryStore struct {
	mu    sync.RWMutex
	store []models.URL
	users []models.User
}

func New() *MemoryStore {
	return &MemoryStore{
		store: make([]models.URL, 0),
		users: make([]models.User, 0),
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

func (m *MemoryStore) Save(url *models.ShortURLRequest) (*models.ShortURLResponse, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	shortCode := GenerateRandomString(6)
	newURL := models.URL{
		ShortCode:   shortCode,
		OriginalURL: url.URL,
		CreatedAt:   time.Now().Unix(),
		UserID:      nil,
	}
	m.store = append(m.store, newURL)
	return &models.ShortURLResponse{ShortURL: shortCode, OriginalURL: url.URL}, nil
}

func (m *MemoryStore) Get(shortCode string) (*models.ShortURLResponse, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	for _, url := range m.store {
		if url.ShortCode == shortCode {
			return &models.ShortURLResponse{ShortURL: "http://localhost:8080/" + shortCode, OriginalURL: url.OriginalURL}, nil
		}
	}
	return nil, errors.New("short URL not found")
}

func (m *MemoryStore) GetAll() ([]models.URL, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	result := make([]models.URL, len(m.store))
	copy(result, m.store)
	for i := 0; i < len(result)-1; i++ {
		for j := i + 1; j < len(result); j++ {
			if result[i].CreatedAt < result[j].CreatedAt {
				result[i], result[j] = result[j], result[i]
			}
		}
	}
	return result, nil
}

func (m *MemoryStore) Health() error {
	return nil // In-memory, always healthy
}