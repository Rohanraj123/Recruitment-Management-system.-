package main

import (
	"net/http"
	"synergy/models" // Adjust the import path based on your project structure
	"synergy/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite" // or another driver depending on your DB
	"gorm.io/gorm"
)

func main() {
	// Initialize the Gin router
	r := gin.Default()

	// Connect to the database
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{}) // Example with SQLite
	if err != nil {
		panic("failed to connect to database")
	}

	// Auto migrate the User model
	if err := db.AutoMigrate(&models.User{}); err != nil {
		panic("failed to migrate database")
	}

	// Define your signup route
	r.POST("/signup", func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Call your create user method
		if err := user.CreateUser(db); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "user created successfully"})
	})

	// Define your login route
	r.POST("/login", func(c *gin.Context) {
		var input struct {
			Email    string `json:"email" binding:"required"`
			Password string `json:"password" binding:"required"`
		}

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		// Fetch user from database by email
		var user models.User
		if err := db.Where("email = ?", input.Email).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			return
		}

		// Compare the provided password with the hashed password
		if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password)); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
			return
		}

		// Generate JWT token
		token, err := utils.GenerateToken(user.Email, user.UserType)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
			return
		}

		// Return the token to the user
		c.JSON(http.StatusOK, gin.H{"token": token})
	})

	// Start the server
	r.Run(":8080")
}
