package main

import (
	"gort/handlers"
	"gort/models"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	app.Use(gin.Logger())

	models.ConnectDB()
	app.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the URL Shortener API",
		})
	})
	app.GET("/:url", handlers.ResolveUrlHandler)
	app.POST("/", handlers.ShortenUrlHandler)

	app.Run(":6969")
}
