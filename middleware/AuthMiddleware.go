package middleware

import (
	"net/http"
	"synergy/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Validate the token
		err := utils.ValidateToken(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token: " + err.Error()})
			c.Abort()
			return
		}

		// Extract the claims after token validation
		claims, err := utils.ExtractClaims(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to extract claims: " + err.Error()})
			c.Abort()
			return
		}

		// Add user info to context
		c.Set("user_type", claims.UserType)
		c.Set("email", claims.Email)

		// Debugging statement to log the user type
		c.Next()
	}
}
