package models

type Profile struct {
	ID             uint `gorm:"primaryKey"`
	ApplicantID    uint
	ResumeFilePath string
	Skills         string
	Education      string
	Experience     string
	Name           string
	Email          string
	Phone          string
}
