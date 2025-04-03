// S:\SDE\Hard Core\Learn\Golang\Projects\URL-Shortner-with-Go\Backend\internal\models\url.go
package models

// ShortURLRequest represents the request body for shortening a URL.
type ShortURLRequest struct {
    URL string `json:"url" binding:"required"`
}

// ShortURLResponse represents the response with the shortened URL.
type ShortURLResponse struct {
    ShortURL    string `json:"short_url"`
    OriginalURL string `json:"original_url"` // Added OriginalURL field
}

// ErrorResponse represents an error response returned by the API.
type ErrorResponse struct {
    Error string `json:"error"`
}

// URLListResponse represents the response for listing all URLs.
type URLListResponse struct {
    ShortCode   string `json:"short_code"`
    OriginalURL string `json:"original_url"`
}