package handlers

import (
	"gort/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type response struct {
	models.URL      `json:"url"`
	XRateRemaining  int           `json:"rate_limit"`
	XRateLimitReset time.Duration `json:"rate_limit_reset"`
}

func ResolveUrlHandler(c *gin.Context) {
	var url models.URL

	if err := models.DB.Where("short_url = ?", c.Param("url")).First(&url).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.Redirect(http.StatusOK, url.LongUrl)
}

func ShortenUrlHandler(c *gin.Context) {
	var url models.URL

	if err := c.ShouldBindJSON(&url); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := url.Sanitize(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	url.ShortUrl = "https://" + uuid.NewString()[:6]
	url.Expiry = time.Second * 60

	models.DB.Create(&url)
	c.JSON(http.StatusOK, gin.H{"data": url})
}
