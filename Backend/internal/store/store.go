
// S:\SDE\Hard Core\Learn\Golang\Projects\URL-Shortner-with-Go\Backend\internal\store\store.go
package store

import "Backend/internal/models"

// Store interface for managing URLs.
type Store interface {
    Save(request *models.ShortURLRequest) (*models.ShortURLResponse, error) // Changed 'url' to 'request'
    Get(shortCode string) (*models.ShortURLResponse, error)
    GetAll() (map[string]string, error)
}