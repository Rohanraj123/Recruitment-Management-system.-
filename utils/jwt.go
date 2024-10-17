package utils

import (
	"errors"
	"strings"
	"synergy/config"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Claims represents the JWT claims
type Claims struct {
	Email    string `json:"email"`
	UserType string `json:"user_type"`
	jwt.StandardClaims
}

// GenerateToken generates a JWT token for a user
func GenerateToken(email, userType string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Email:    email,
		UserType: userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.GetJWTSecret())
}

// ValidateToken validates the JWT token and returns error if invalid
func ValidateToken(c *gin.Context) error {
	// Get the Authorization header
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		return errors.New("authorization token not provided")
	}

	// Extract the token from the header
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	// Parse the token and validate it
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return config.GetJWTSecret(), nil
	})

	if err != nil {
		return errors.New("error parsing token: " + err.Error())
	}

	// Check token validity
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		if claims.ExpiresAt < time.Now().Unix() {
			return errors.New("token expired")
		}
	} else {
		return errors.New("invalid token")
	}

	return nil
}

// ExtractClaims extracts and returns the claims from the JWT token
func ExtractClaims(c *gin.Context) (*Claims, error) {
	// Validate the token first
	err := ValidateToken(c)
	if err != nil {
		return nil, err
	}

	// Get the token from the header
	tokenString := c.GetHeader("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	// Parse the token and extract claims
	token, _ := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return config.GetJWTSecret(), nil
	})

	// Return the claims if the token is valid
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
