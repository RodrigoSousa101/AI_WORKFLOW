package taskuser

import (
	"net/http"

	"github.com/RodrigoSousa101/ai_workflow/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func getAllTaskUser(c *gin.Context) {
	var taskUsers []models.TaskUser
	db := c.MustGet("db").(*gorm.DB)

	if err := db.Find(&taskUsers).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Failed to retrieve Task Users"})
		return
	}

	if len(taskUsers) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No Task Users found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"TaskUsers": taskUsers})

}
