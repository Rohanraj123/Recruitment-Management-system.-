package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string `json:"name"`
	Email        string `json:"email" gorm:"unique"`
	PasswordHash string `json:"-"`
	UserType     string `json:"user_type"` // "Admin" or "Applicant"
	Headline     string `json:"headline"`
	Address      string `json:"address"`
	Password     string `json:"password"` // Add this line
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
