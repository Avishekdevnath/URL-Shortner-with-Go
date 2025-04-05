// S:\SDE\Hard Core\Learn\Golang\Projects\URL-Shortner-with-Go\Backend\internal\store\postgres\postgres.go
package postgres

import (
	"fmt"
	"time"
	"Backend/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresStore struct {
	db *gorm.DB
}

func New(dsn string) (*PostgresStore, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	// Explicitly check the 'public' schema for tables
	var urlTableExists, userTableExists int
	db.Raw("SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = 'public' AND table_name = 'urls'").Scan(&urlTableExists)
	db.Raw("SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = 'public' AND table_name = 'users'").Scan(&userTableExists)

	// Only migrate if tables don’t exist
	if urlTableExists == 0 {
		if err := db.AutoMigrate(&models.URL{}); err != nil {
			return nil, fmt.Errorf("failed to migrate URLs table: %v", err)
		}
	}
	if userTableExists == 0 {
		if err := db.AutoMigrate(&models.User{}); err != nil {
			return nil, fmt.Errorf("failed to migrate Users table: %v", err)
		}
	}

	return &PostgresStore{db: db}, nil
}

// Save saves the URL in PostgreSQL.
func (p *PostgresStore) Save(url *models.ShortURLRequest) (*models.ShortURLResponse, error) {
	shortCode := fmt.Sprintf("%d", time.Now().Unix())
	newURL := models.URL{
		ShortCode:   shortCode,
		OriginalURL: url.URL,
		CreatedAt:   time.Now().Unix(),
	}

	if err := p.db.Create(&newURL).Error; err != nil {
		return nil, fmt.Errorf("failed to save URL in PostgreSQL: %v", err)
	}

	return &models.ShortURLResponse{ShortURL: shortCode, OriginalURL: url.URL}, nil
}

// Get retrieves the URL from PostgreSQL by short code.
func (p *PostgresStore) Get(shortCode string) (*models.ShortURLResponse, error) {
	var url models.URL
	if err := p.db.Where("short_code = ?", shortCode).First(&url).Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve URL from PostgreSQL: %v", err)
	}
	return &models.ShortURLResponse{ShortURL: shortCode, OriginalURL: url.OriginalURL}, nil
}

// GetAll retrieves all shortened URLs.
func (p *PostgresStore) GetAll() ([]models.URL, error) {
	var urls []models.URL
	if err := p.db.Order("created_at desc").Find(&urls).Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve all URLs: %v", err)
	}
	return urls, nil
}

// Health checks if PostgreSQL is responsive.
func (p *PostgresStore) Health() error {
	return p.db.Raw("SELECT 1").Error
}