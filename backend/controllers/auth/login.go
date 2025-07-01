package auth

import (
	"net/http"

	"github.com/RodrigoSousa101/ai_workflow/models"

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
	if err := db.Where("email = ? AND password = ?", user.Email, user.Password).First(&existingUser).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}
	existingUser.Password = "" // Clear password before sending response
	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "user": existingUser})
}
