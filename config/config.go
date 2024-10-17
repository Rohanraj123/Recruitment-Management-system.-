package config

import (
	"synergy/models"

	"gorm.io/driver/sqlite" // Use your specific driver
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database!")
	}

	// Auto migrate the Job model to create the jobs table
	if err := DB.AutoMigrate(&models.Job{}, &models.Profile{}); err != nil {
		panic("Failed to migrate database!")
	}
}

var JWTSecret = []byte("S7BcXcSi580kdm3L9n1Cyy+53woz8wMgQVevAjNS9xA=")

func GetJWTSecret() []byte {
	return JWTSecret
}
