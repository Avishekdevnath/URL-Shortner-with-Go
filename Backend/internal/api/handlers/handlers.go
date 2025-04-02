package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"Backend/internal/models"
	"Backend/internal/service"  // Local import for the service layer
)

// ShortenURL handles POST requests to shorten URLs.
func ShortenURL(c *gin.Context, urlService *service.URLService) {
	var request models.ShortURLRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid input"})
		return
	}

	// Call the service to shorten the URL
	response, err := urlService.ShortenURL(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

// RedirectURL handles GET requests to redirect to the original URL.
func RedirectURL(c *gin.Context, urlService *service.URLService) {
	shortCode := c.Param("code")

	// Call the service to retrieve the original URL
	response, err := urlService.GetOriginalURL(shortCode)
	if err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "URL not found"})
		return
	}

	c.Redirect(http.StatusMovedPermanently, response.ShortURL)
}
