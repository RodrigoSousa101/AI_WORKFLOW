package task

import (
	"github.com/RodrigoSousa101/ai_workflow/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteTask(c *gin.Context) {
	var task models.Task
	TaskID := c.Param("id")
	db := c.MustGet("db").(*gorm.DB)

	if db.First(&task, "id = ?", TaskID).Error != nil {
		c.JSON(404, gin.H{"error": "Task not found"})
		return
	}
	if err := db.Delete(&task, "id = ?", TaskID).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete task"})
		return
	}

	c.JSON(200, gin.H{"message": "Task deleted successfully", "Task": task})
}
