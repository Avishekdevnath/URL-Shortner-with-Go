package store

import "github.com/Avishekdevnath/URL-Shortner-with-Go/Backend/internal/models"

// Store interface for managing URLs.
type Store interface {
    Save(url *models.ShortURLRequest) (*models.ShortURLResponse, error)
    Get(shortCode string) (*models.ShortURLResponse, error)
}
