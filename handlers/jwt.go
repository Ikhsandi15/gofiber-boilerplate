package handlers

import (
	"gofiber_boilerplate/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateJwtToken(user *models.User) (string, error) {
	var jwtKey = []byte(os.Getenv("JWT_SECRET"))

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(10 * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
