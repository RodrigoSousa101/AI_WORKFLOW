package users

import (
	"github.com/RodrigoSousa101/ai_workflow/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteUser(c *gin.Context) {
	var User models.User
	userID := c.Param("id")
	db := c.MustGet("db").(*gorm.DB)
	if db.First(&User, "id = ?", userID).Error != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}
	if err := db.Delete(&User, "id = ?", userID).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(200, gin.H{"message": "User deleted successfully", "user": User})
}
