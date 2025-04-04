// S:\SDE\Hard Core\Learn\Golang\Projects\URL-Shortner-with-Go\Backend\internal\api\handlers\handlers.go

package handlers

import (
	"Backend/internal/models"
	"Backend/internal/service"
	"fmt"
	"log"
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

	log.Printf("Shortened URL: %s", response.ShortURL)
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

	response, err := urlService.GetOriginalURL(shortCode)
	if err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "URL not found"})
		return
	}

	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET")
	fmt.Println("Redirecting to", response.OriginalURL)
	c.Redirect(http.StatusMovedPermanently, response.OriginalURL)
}

// GetAllURLs handles GET requests to list all shortened URLs.
// @Summary List all shortened URLs
// @Description Returns a list of all shortened URLs with their original URLs
// @Tags URL
// @Produce json
// @Success 200 {array} models.URL "List of shortened URLs with metadata"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /urls [get]
func GetAllURLs(c *gin.Context, urlService *service.URLService) {
	allURLs, err := urlService.GetAllURLs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}

	// Return the full URL structs (includes UserID and CreatedAt)
	c.JSON(http.StatusOK, allURLs)
}