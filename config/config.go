package config

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB
var JWTSecret = []byte("S7BcXcSi580kdm3L9n1Cyy+53woz8wMgQVevAjNS9xA=")

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("recruitment.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database", err)
	}
	log.Println("Database connected!")
}

func GetJWTSecret() []byte {
	return JWTSecret
}
