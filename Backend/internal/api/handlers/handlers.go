package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Avishekdevnath/URL-Shortner-with-Go/Backend/internal/models"
	"github.com/Avishekdevnath/URL-Shortner-with-Go/Backend/internal/store/memory"
)

// Initialize memory store
var store = memory.New()

// ShortenURL handles POST requests to shorten URLs.
func ShortenURL(c *gin.Context) {
	var request models.ShortURLRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid input"})
		return
	}

	// Save the URL and get the short code
	response, err := store.Save(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Error saving URL"})
		return
	}

	c.JSON(http.StatusOK, response)
}

// RedirectURL handles GET requests to redirect to the original URL.
func RedirectURL(c *gin.Context) {
	shortCode := c.Param("code")

	// Get the original URL for the short code
	response, err := store.Get(shortCode)
	if err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "URL not found"})
		return
	}

	c.Redirect(http.StatusMovedPermanently, response.ShortURL)
}
