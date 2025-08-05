package auth

import (
	"fmt"
	"net/http"
	"os"

	"github.com/RodrigoSousa101/ai_workflow/models"
	"github.com/RodrigoSousa101/ai_workflow/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

func Refresh(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	fmt.Println("Refresh token cookie:", refreshToken, "Error:", err)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "no refresh token"})
		return
	}

	// Verificar validade
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("REFRESH_SECRET")), nil
	})

	if err != nil || !token.Valid {
		fmt.Println("Refresh token inválido ou erro:", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid refresh token"})
		return
	}
	fmt.Println("Refresh token válido")

	claims := token.Claims.(jwt.MapClaims)
	userIdStr, ok := claims["sub"].(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid subject in token"})
		return
	}

	db := c.MustGet("db").(*gorm.DB)
	var user models.User
	if err := db.First(&user, "id = ?", userIdStr).Error; err != nil {
		fmt.Println("Usuário não encontrado:", userIdStr)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}
	fmt.Println("Usuário encontrado:", user.ID)

	// Criar novo access token
	newAccessToken, err := utils.CreateAccessToken(user)
	if err != nil {
		fmt.Println("Erro ao criar novo access token:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create new access token"})
		return
	}

	fmt.Println("Novo access token criado com sucesso")
	c.JSON(http.StatusOK, gin.H{"access": newAccessToken})
}
