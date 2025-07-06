package utils

import (
	"os"
	"time"

	"github.com/RodrigoSousa101/ai_workflow/models"
	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(user models.User) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  user.ID,
		"role": user.Role,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	})
	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))
}
