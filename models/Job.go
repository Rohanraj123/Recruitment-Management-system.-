package models

import (
	"time"

	"gorm.io/gorm"
)

type Job struct {
	gorm.Model
	Title             string
	Description       string
	PostedOn          time.Time
	TotalApplications int
	CompanyName       string
	PostedBy          uint
}
