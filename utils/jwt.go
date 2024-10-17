package utils

import (
	"synergy/config"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(email string, userType string) (string, error) {
	claims := jwt.MapClaims{
		"email":     email,
		"user_type": userType,
		"exp":       time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.GetJWTSecret())
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return config.GetJWTSecret(), nil
	})
}
