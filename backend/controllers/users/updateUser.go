package users

import (
	"net/http"
	"time"

	"github.com/RodrigoSousa101/ai_workflow/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateUser(c *gin.Context) {
	var UserUpdate models.User
	var User models.User
	UserID := c.Param("id")

	if err := c.ShouldBindJSON((&UserUpdate)); err != nil {
		c.JSON((http.StatusBadRequest), gin.H{"error": err.Error()})
	}

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("id = ?", UserID).First(&User).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := db.Model(&User).Where("id = ?", UserID).Updates(UserUpdate).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}
	User.UpdatedAt = time.Now()
	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully", "user": User})
}
