package task

import (
	"net/http"

	"github.com/RodrigoSousa101/ai_workflow/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetTask(c *gin.Context) {
	var task models.Task
	taskID := c.Param("id")
	db := c.MustGet("db").(*gorm.DB)

	if err := db.Find(&task, "id = ?", taskID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found with this id"})
	}

	c.JSON(http.StatusFound, gin.H{"message": "Task found", "Task:": task})
}
