package service

import (
    "errors"
    "Backend/internal/models" // Local import for models
    "Backend/internal/store"  // Local import for store interface
)

// URLService contains the business logic for the URL shortener.
type URLService struct {
    store store.Store
}

// NewURLService creates a new instance of URLService.
func NewURLService(store store.Store) *URLService {
    return &URLService{store: store}
}

// ShortenURL generates a shortened URL and saves it in the store.
func (s *URLService) ShortenURL(request *models.ShortURLRequest) (*models.ShortURLResponse, error) {
    // Validate the input URL
    if request.URL == "" {
        return nil, errors.New("URL is required")
    }

    // Save the URL and get the shortened version
    response, err := s.store.Save(request)
    if err != nil {
        return nil, err
    }

    return response, nil
}

// GetOriginalURL retrieves the original URL for a given short code.
func (s *URLService) GetOriginalURL(shortCode string) (*models.ShortURLResponse, error) {
    // Retrieve the original URL using the short code
    response, err := s.store.Get(shortCode)
    if err != nil {
        return nil, err
    }

    return response, nil
}

// GetAllURLs retrieves all stored URLs.
func (s *URLService) GetAllURLs() (map[string]string, error) {
    // Call the store's GetAll method and handle the error
    urls, err := s.store.GetAll()
    if err != nil {
        return nil, err
    }

    return urls, nil
}