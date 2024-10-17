package main

import (
	"synergy/controllers"

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

	// Start the server
	r.Run(":8080")
}
