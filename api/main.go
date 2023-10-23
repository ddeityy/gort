package main

import (
	"gort/handlers"
	"gort/models"

	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	models.ConnectDB()

	app.GET("/:url", logger.SetLogger(), handlers.ResolveUrlHandler)
	app.POST("/", logger.SetLogger(), handlers.ShortenUrlHandler)

	app.Run("localhost:6969")
}
