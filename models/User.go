package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID              uint `gorm:"primaryKey"`
	Name            string
	Email           string `gorm:"unique"`
	Address         string
	UserType        string // "APPLICANT" or "ADMIN"
	PasswordHash    string
	ProfileHeadline string
	Profile         Profile
}

func (u *User) CreateUser(db *gorm.DB) error {
	// Hash the password before saving
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.PasswordHash = string(hashedPassword) // Set the hashed password
	return db.Create(u).Error
}

func GetUserByEmail(db *gorm.DB, email string) (*User, error) {
	var user User
	err := db.Where("email = ?", email).First(&user).Error
	return &user, err
}
