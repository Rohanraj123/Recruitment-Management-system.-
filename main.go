package main

import (
	"synergy/controllers"
	"synergy/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the Gin router
	r := gin.Default()

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
	}

	// Start the server
	r.Run(":8080")
}
