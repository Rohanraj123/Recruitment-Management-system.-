package controllers

import (
	"log" // Import the log package
	"net/http"
	"synergy/config"
	"synergy/models"
	"synergy/utils"

	"github.com/gin-gonic/gin"
)

func CreateJob(c *gin.Context) {
	// Extract the JWT claims to check user role
	userClaims, err := utils.ExtractClaims(c)
	if err != nil || userClaims.UserType != "Admin" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Only Admin users are allowed to create jobs"})
		return
	}

	// Define the structure of the incoming job request
	var job models.Job
	if err := c.ShouldBindJSON(&job); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Validate required fields
	if job.Title == "" || job.Description == "" || job.CompanyName == "" || job.Location == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title, Description, Company Name, and Location are required fields."})
		return
	}

	// Create the job in the database
	if err := config.DB.Create(&job).Error; err != nil {
		// Log the error for debugging purposes
		log.Printf("Error creating job: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create job"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Job created successfully",
		"job":     job,
	})
}
