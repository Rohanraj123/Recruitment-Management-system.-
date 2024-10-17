package controllers

import (
	"net/http"
	"synergy/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.PasswordHash), 8)
	input.PasswordHash = string(hashedPassword)

	// You have to create DB model in models package for database config.
	models.DB.create(&input)

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully!"})
}
