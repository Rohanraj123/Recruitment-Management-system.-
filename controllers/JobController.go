package controllers

import (
	"log"
	"net/http"
	"synergy/config"
	"synergy/models"
	"synergy/utils"

	"github.com/gorilla/mux"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"encoding/json"
)

var db *gorm.DB

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
	if job.Title == "" || job.Description == "" || job.CompanyName == "" {
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

// Get Job by ID handler
func getJobByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	jobID := vars["job_id"]

	var job models.Job
	if result := db.First(&job, jobID); result.Error != nil {
		http.Error(w, "Job not found", http.StatusNotFound)
		return
	}

	var applicants []models.User
	db.Where("profile.applicant_id = ?", jobID).Find(&applicants)

	response := struct {
		Job        models.Job    `json:"job"`
		Applicants []models.User `json:"applicants"`
	}{
		Job:        job,
		Applicants: applicants,
	}
	json.NewEncoder(w).Encode(response)
}

// Get all applicants handler
func getAllApplicantsHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.Where("user_type = ?", "APPLICANT").Find(&users)
	json.NewEncoder(w).Encode(users)
}

// Get Applicant by ID handler
func getApplicantByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	applicantID := vars["applicant_id"]

	var profile models.Profile
	if result := db.First(&profile, applicantID); result.Error != nil {
		http.Error(w, "Applicant not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(profile)
}

// Get all jobs handler
func getAllJobsHandler(w http.ResponseWriter) {
	var jobs []models.Job
	db.Find(&jobs)
	json.NewEncoder(w).Encode(jobs)
}

// Apply for Job handler
func applyForJobHandler(w http.ResponseWriter) {
	// Check if the user is an applicant, then process the application
	json.NewEncoder(w).Encode("Application submitted successfully")
}
