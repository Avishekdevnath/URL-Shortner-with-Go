package handlers

import (
	"Backend/internal/models"
	"Backend/internal/service" // Local import for the service layer
	"fmt"
	"log" // For logging the shortened URL
	"net/http"

	"github.com/gin-gonic/gin"
)

// ShortenURL handles POST requests to shorten URLs.
// @Summary Shorten a URL
// @Description Takes a long URL and returns a shortened version
// @Tags URL
// @Accept json
// @Produce json
// @Param request body models.ShortURLRequest true "URL to shorten"
// @Success 200 {object} models.ShortURLResponse "Shortened URL"
// @Failure 400 {object} models.ErrorResponse "Invalid input"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /api/shorten [post]
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

	// Log the shortened URL to the console
	log.Printf("Shortened URL: %s", response.ShortURL) // This will print the shortened URL in the console

	c.JSON(http.StatusOK, response)
}

// RedirectURL handles GET requests to redirect to the original URL.
// @Summary Redirect to original URL
// @Description Redirects to the original URL based on the short code
// @Tags URL
// @Produce plain
// @Param code path string true "Short URL code"
// @Success 301 "Redirects to the original URL"
// @Failure 404 {object} models.ErrorResponse "URL not found"
// @Router /{code} [get]
func RedirectURL(c *gin.Context, urlService *service.URLService) {
	shortCode := c.Param("code")

	// Call the service to retrieve the original URL
	response, err := urlService.GetOriginalURL(shortCode)
	if err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "URL not found"})
		return
	}

	// Add CORS headers to allow browser requests
	c.Header("Access-Control-Allow-Origin", "*") // Allow all origins for testing
	c.Header("Access-Control-Allow-Methods", "GET")

	fmt.Println("line 44", response.OriginalURL)
	c.Redirect(http.StatusMovedPermanently, response.OriginalURL)
}

// GetAllURLs handles GET requests to list all shortened URLs.
// @Summary List all shortened URLs
// @Description Returns a list of all shortened URLs with their original URLs
// @Tags URL
// @Produce json
// @Success 200 {array} models.URLListResponse "List of shortened URLs"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /urls [get]
func GetAllURLs(c *gin.Context, urlService *service.URLService) {
	// Fetch all URLs from the service layer
	allURLs, err := urlService.GetAllURLs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}

	// Prepare the response
	var response []models.URLListResponse
	for shortCode, originalURL := range allURLs {
		response = append(response, models.URLListResponse{
			ShortCode:   shortCode,
			OriginalURL: originalURL,
		})
	}

	// Return the list of URLs in the response
	c.JSON(http.StatusOK, response)
}