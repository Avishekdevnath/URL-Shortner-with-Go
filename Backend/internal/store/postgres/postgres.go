package postgres

import (
	"errors"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"Backend/internal/models" // Local import for models
	"Backend/internal/store"  // Local import for store interface
)

// PostgresStore represents a PostgreSQL-based store.
type PostgresStore struct {
	db *gorm.DB
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
	// Generate a unique short code (could be improved with a better generation strategy)
	shortCode := fmt.Sprintf("short%d", len(url.URL))

	// Create a new ShortURLResponse
	shortURL := models.ShortURLResponse{
		ShortURL: shortCode,
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
