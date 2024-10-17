package main

import (
	"synergy/config"
	"synergy/controllers"
	"synergy/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the Gin router
	r := gin.Default()

	config.ConnectDatabase()

	// SignUp endpoint
	r.POST("/signup", func(c *gin.Context) {
		controllers.SignUp(c)
	})

	// LogIn endpoint
	r.POST("/login", func(c *gin.Context) {
		controllers.LogIn(c)
	})

	// Protected routes with authentication
	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.POST("/uploadResume", controllers.UploadResume)

		// Admin-specific routes
		admin := auth.Group("/admin")
		admin.Use(middleware.AuthMiddleware())
		{
			admin.POST("/job", controllers.CreateJob)

			// Fetch job details by job_id
			// admin.GET("/job/:job_id", controllers.GetJobDetails)

			// Fetch all applicants
			// admin.GET("/applicants", controllers.GetAllApplicants)

			// Fetch a specific applicant by applicant_id
			// admin.GET("/applicant/:applicant_id", controllers.GetApplicantDetails)

		}
	}

	// General route for fetching job openings
	// auth.GET("/jobs", controllers.GetJobs)

	// Apply for a job (only for Applicants)
	// auth.POST("/jobs/apply", controllers.ApplyForJob)

	// Start the server
	r.Run(":8080")
}
