package users

import (
	"github.com/RodrigoSousa101/ai_workflow/models"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	var User models.User
	userID := c.Param("id")

	db := c.MustGet("db").(*gorm.DB)
	if err := db.First(&User, "id = ?", userID).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	c.JSON(200, gin.H{"message": "User found", "user": User})
}
