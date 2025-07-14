package task

import (
	"net/http"
	"time"

	"github.com/RodrigoSousa101/ai_workflow/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateTask(c *gin.Context) {
	var task models.Task
	var existingTask models.Task
	TaskID := c.Param("id")
	db := c.MustGet("db").(*gorm.DB)

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.First(&existingTask, "id = ?", TaskID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error:": "Task not found"})
		return
	}

	if err := db.Model(&existingTask).Updates(task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update Task"})
		return
	}

	existingTask.Updated_at = time.Now()
	c.JSON(http.StatusOK, gin.H{"message:": "task updated sucessfuly", "Task:": existingTask})
}
