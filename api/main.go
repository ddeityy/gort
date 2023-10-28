package main

import (
	"gort/handlers"
	"gort/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	app.Use(gin.Logger())
	app.Use(cors.Default())

	models.ConnectDB()
	app.GET("/:url", handlers.ResolveUrl)
	app.POST("/", handlers.ShortenUrl)

	app.Run(":6969")
}
