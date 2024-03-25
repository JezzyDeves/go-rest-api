package db

import (
	"github.com/JezzyDeves/go-rest-api/db/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Gets the DB connection and performs migrations
func getDBConnection() *gorm.DB {
	db, err := gorm.Open(
		sqlite.Open("file::memory:?cache=shared"),
		&gorm.Config{})

	if err != nil {
		panic("Issue connecting to the database")
	}

	db.AutoMigrate(&models.Employee{})

	return db
}

var Database *gorm.DB = getDBConnection()
