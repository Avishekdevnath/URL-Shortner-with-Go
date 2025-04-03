// S:\SDE\Hard Core\Learn\Golang\Projects\URL-Shortner-with-Go\Backend\internal\store\postgres\postgres.go

package postgres

import (
	"errors"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"Backend/internal/models" // Local import for models
	"Backend/internal/store"  // Local import for store interface
	"math/rand"
	"time"
)

// PostgresStore represents a PostgreSQL-based store.
type PostgresStore struct {
	db *gorm.DB
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

// GetAll retrieves all stored short URLs.
func (p *PostgresStore) GetAll() (map[string]string, error) {
	var shortURLs []models.ShortURLResponse
	result := make(map[string]string)

	// Query the database for all short URLs
	if err := p.db.Find(&shortURLs).Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve all URLs: %v", err)
	}

	// Convert the result to a map
	for _, shortURL := range shortURLs {
		result[shortURL.ShortURL] = shortURL.OriginalURL // Changed to shortURL.ShortURL
	}

	return result, nil
}

// New creates a new PostgresStore and connects to the database.
func New(dsn string) (*PostgresStore, error) {
	// Open the connection to PostgreSQL using GORM
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	// Migrate the schema (create table if not exists)
	err = db.AutoMigrate(&models.ShortURLRequest{}, &models.ShortURLResponse{})
	if err != nil {
		return nil, fmt.Errorf("failed to migrate database: %v", err)
	}

	return &PostgresStore{db: db}, nil
}

// Save stores the original URL and returns a shortened URL.
func (p *PostgresStore) Save(url *models.ShortURLRequest) (*models.ShortURLResponse, error) {
	// Generate a random short code
	shortCode := GenerateRandomString(6)

	// Create a new ShortURLResponse
	shortURL := models.ShortURLResponse{
		ShortURL:    shortCode,
		OriginalURL: url.URL,
	}

	// Save to the database
	if err := p.db.Create(&shortURL).Error; err != nil {
		return nil, fmt.Errorf("failed to save URL: %v", err)
	}

	return &shortURL, nil
}

// Get retrieves the original URL for the given short code.
func (p *PostgresStore) Get(shortCode string) (*models.ShortURLResponse, error) {
	var shortURL models.ShortURLResponse

	// Query the database for the short URL
	if err := p.db.Where("short_url = ?", shortCode).First(&shortURL).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("short URL not found")
		}
		return nil, fmt.Errorf("failed to retrieve URL: %v", err)
	}

	return &shortURL, nil
}

// Ensure that PostgresStore implements the Store interface
var _ store.Store = &PostgresStore{}