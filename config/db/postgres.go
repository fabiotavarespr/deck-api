package db

import (
	"deck-api/models"

	"github.com/jinzhu/gorm"
)

// SetupDB configure a new connection using gorm.
// TODO: Include env vars.
func SetupDB() (*gorm.DB, error) {
	db, err := gorm.Open("postgres", "host=database port=5432 user=postgres dbname=deck password=postgres sslmode=disable")
	if err != nil {
		return nil, err
	}

	db.LogMode(true)

	// Auto migrate is enabled just to accomplish this challenge. In a real situation, migrations should be properly provided.
	if err := db.AutoMigrate(&models.Deck{}).Error; err != nil {
		return nil, err
	}

	return db, nil
}
