package utils

import (
	"os"
	"time"

	"github.com/RodrigoSousa101/ai_workflow/models"
	"github.com/golang-jwt/jwt/v5"
)

func CreateAccessToken(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  user.ID,
		"role": user.Role,
		"exp":  time.Now().Add(15 * time.Minute).Unix(),
	})
	return token.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
}

func CreateRefreshToken(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(7 * 24 * time.Hour).Unix(),
	})
	return token.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
}
