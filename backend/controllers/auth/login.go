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

	existingUser.Password = "" // Clear password before sending response
	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "user": existingUser})
}
