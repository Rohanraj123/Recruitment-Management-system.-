package controllers

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func UploadResume(c *gin.Context) {
	// Check if the user is an Applicant
	userType, exists := c.Get("user_type")
	if !exists || userType != "Applicant" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only applicants can upload resumes"})
		return
	}

	// Get the file from the request
	file, err := c.FormFile("resume")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}

	// Validate file type (PDF or DOCX)
	ext := filepath.Ext(file.Filename)
	if ext != ".pdf" && ext != ".docx" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file type. Only PDF or DOCX allowed"})
		return
	}

	// Define file storage path
	storagePath := filepath.Join("uploads", file.Filename)

	// Save the uploaded file
	if err := c.SaveUploadedFile(file, storagePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Resume uploaded successfully", "file_path": storagePath})
}
