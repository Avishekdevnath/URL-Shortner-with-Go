// S:\SDE\Hard Core\Learn\Golang\Projects\URL-Shortner-with-Go\Backend\internal\models\url.go




package models

// User represents an optional user profile for authenticated users.
type User struct {
    ID        string `json:"id"` // Unique ID from OAuth (e.g., Google sub)
    Email     string `json:"email"`
    Name      string `json:"name"`
    CreatedAt int64  `json:"created_at"` // Unix timestamp
}

// URL represents a stored URL with metadata, optionally linked to a user.
type URL struct {
    ShortCode   string  `json:"short_code"`
    OriginalURL string  `json:"original_url"`
    CreatedAt   int64   `json:"created_at"` // Unix timestamp for sorting
    UserID      *string `json:"user_id,omitempty"` // Optional: null for anonymous, set for logged-in
}

// ShortURLRequest represents the request body for shortening a URL.
type ShortURLRequest struct {
    URL string `json:"url" binding:"required"`
}

// ShortURLResponse represents the response with the shortened URL.
type ShortURLResponse struct {
    ShortURL    string `json:"short_url"`
    OriginalURL string `json:"original_url"`
}

// ErrorResponse represents an error response returned by the API.
type ErrorResponse struct {
    Error string `json:"error"`
}

// URLListResponse represents the response for listing all URLs.
// Kept for frontend compatibility.
type URLListResponse struct {
    ShortCode   string `json:"short_code"`
    OriginalURL string `json:"original_url"`
}