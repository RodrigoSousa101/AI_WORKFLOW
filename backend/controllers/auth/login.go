package auth

import (
	"net/http"

	"github.com/RodrigoSousa101/ai_workflow/models"
	"github.com/RodrigoSousa101/ai_workflow/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.MustGet("db").(*gorm.DB)
	var existingUser models.User

	if err := db.Where("email = ?", user.Email).First(&existingUser).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}
	if !utils.CheckPasswordHash(user.Password, existingUser.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid credentials"})
		return
	}

	accessToken, err := utils.CreateAccessToken(existingUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create access token"})
		return
	}

	refreshToken, err := utils.CreateRefreshToken(existingUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create refresh token"})
		return
	}

	// Enviar refresh token como cookie HttpOnly
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie(
		"refresh_token",
		refreshToken,
		60*60*24*7,
		"/",
		"",
		false, // HTTP local, Secure false
		true,  // HttpOnly
	)

	// O access token vai no corpo da resposta
	existingUser.Password = ""
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"user":    existingUser,
		"access":  accessToken,
	})
}
