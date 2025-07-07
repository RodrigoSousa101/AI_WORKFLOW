package task

import (
	"net/http"

	"github.com/RodrigoSousa101/ai_workflow/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllTasks(c *gin.Context) {
	var task []models.Task
	db := c.MustGet("db").(*gorm.DB)

	if err := db.Find(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retriev tasks"})
		return
	}

	if len(task) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No tasks found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Tasks": task})
}
