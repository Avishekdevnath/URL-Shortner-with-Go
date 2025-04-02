package models

// ShortURLRequest represents the request body for shortening a URL.
type ShortURLRequest struct {
    URL string `json:"url" binding:"required"`
}

// ShortURLResponse represents the response with the shortened URL.
type ShortURLResponse struct {
    ShortURL string `json:"short_url"`
}

// ErrorResponse represents an error response returned by the API.
type ErrorResponse struct {
    Error string `json:"error"`
}