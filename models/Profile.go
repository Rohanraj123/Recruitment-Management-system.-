package models

import "gorm.io/gorm"

type Profile struct {
	gorm.Model
	UserId     uint
	ResumeFile string
	Skills     string
	Education  string
	Experience string
	Name       string
	Email      string
	Phone      string
}
