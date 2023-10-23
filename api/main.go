package main

import (
	"gort/handlers"
	"gort/models"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	models.ConnectDB()

	app.Use(gin.Logger())
	app.Use(gin.Recovery())

	app.GET("/:url", handlers.ResolveUrlHandler)
	app.POST("/", handlers.ShortenUrlHandler)

	app.Run("localhost:6969")
	log.Println("Server started.")
}
