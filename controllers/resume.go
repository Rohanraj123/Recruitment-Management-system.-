package controllers

import (
	"net/http"
	"synergy/services"

	"github.com/gin-gonic/gin"
)

func UploadResume(c *gin.Context) {
	file, _ := c.FormFile("resume")

	// Check file type (PDF/DOCX)
	if file.Header.Get("Content-Type") != "application/pdf" && file.Header.Get("Content-Type") != "application/vnd.openxmlformats-officedocument.wordpressingml.document" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Only PDF and DOCX formats are allowed"})
		return
	}

	// Save file
	filePath := "./uploads/" + file.Filename
	c.SaveUploadedFile(file, filePath)

	// Call the resume parsing API
	parsedData, err := services.ParseResume(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Resume parsing failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"parsedData": string(parsedData)})
}
