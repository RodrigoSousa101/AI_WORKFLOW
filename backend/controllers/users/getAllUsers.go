package users

import (
	"net/http"

	"github.com/RodrigoSousa101/ai_workflow/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllUsers(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var users []models.User

	if err := db.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}

	if len(users) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Users not found"})
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}
