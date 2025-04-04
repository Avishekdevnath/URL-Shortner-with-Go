// S:\SDE\Hard Core\Learn\Golang\Projects\URL-Shortner-with-Go\Backend\internal\service\service.go

package service

import (
	"errors"
	"Backend/internal/models"
	"Backend/internal/store" // Added back the import
)

// URLService contains the business logic for the URL shortener.
type URLService struct {
	store store.Store // Explicitly use store.Store
}

// NewURLService creates a new instance of URLService.
func NewURLService(store store.Store) *URLService {
	return &URLService{store: store}
}

// ShortenURL generates a shortened URL and saves it in the store.
func (s *URLService) ShortenURL(request *models.ShortURLRequest) (*models.ShortURLResponse, error) {
	if request.URL == "" {
		return nil, errors.New("URL is required")
	}
	response, err := s.store.Save(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// GetOriginalURL retrieves the original URL for a given short code.
func (s *URLService) GetOriginalURL(shortCode string) (*models.ShortURLResponse, error) {
	response, err := s.store.Get(shortCode)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// GetAllURLs retrieves all stored URLs.
func (s *URLService) GetAllURLs() ([]models.URL, error) {
	urls, err := s.store.GetAll()
	if err != nil {
		return nil, err
	}
	return urls, nil
}