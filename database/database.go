package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"taskgo/models"
)

func Instance() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("database/database.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database.")
	}

	db.AutoMigrate(&models.Task{})

	return db
}
