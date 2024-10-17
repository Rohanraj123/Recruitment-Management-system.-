package models

import "time"

type Job struct {
	ID                uint `gorm:"primaryKey"`
	Title             string
	Description       string
	PostedOn          time.Time
	TotalApplications int
	CompanyName       string
	PostedBy          User
}
