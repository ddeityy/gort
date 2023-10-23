package models

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	database, err := gorm.Open(sqlite.Open("gort.sqlite"), &gorm.Config{})

	if err != nil {
		log.Panic("Could not connect to database.")
	}

	err = database.AutoMigrate(&URL{})
	if err != nil {
		return
	}

	DB = database
}
