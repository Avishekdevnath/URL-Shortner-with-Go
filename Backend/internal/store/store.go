package store

import "Backend/internal/models"

// Store interface for managing URLs.
type Store interface {
    Save(url *models.ShortURLRequest) (*models.ShortURLResponse, error)
    Get(shortCode string) (*models.ShortURLResponse, error)
}
