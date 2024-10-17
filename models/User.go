package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name         string
	Email        string
	Address      string
	PasswordHash string
	UserType     string
	Profile      string
}
