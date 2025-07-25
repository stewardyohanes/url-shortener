package handlers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stewardyohanes/url-shortener/models"
	"github.com/stewardyohanes/url-shortener/services"
)

func Shorten(c *gin.Context) {
	var req models.ShortenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortCode := uuid.New().String()[:8]

	newURL := models.URL{
		OriginalURL: req.URL,
		ShortCode:   shortCode,
		CreatedAt:   time.Now(),
	}

	if err := services.CreateURL(newURL); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create short URL"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"short_url": fmt.Sprintf("%s/%s", os.Getenv("BASE_URL"), shortCode)})
}

func Redirect(c *gin.Context) {
	shortCode := c.Param("short_code")

	url, err := services.GetURLByShortCode(shortCode)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
	}

	if err := services.UpdateURLVisitCount(shortCode); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update visit count"})
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, url.OriginalURL)
}
