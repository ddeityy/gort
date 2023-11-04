package main

import (
	"gort/handlers"
	"gort/models"
	"log"

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

	err := app.Run(":6969")
	if err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
