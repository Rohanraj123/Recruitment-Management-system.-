package models

import "gorm.io/gorm"

// Job struct defines the Job model
type Job struct {
	gorm.Model
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	CompanyName string `json:"company_name" binding:"required"`
	Location    string `json:"location" binding:"required"`
}
