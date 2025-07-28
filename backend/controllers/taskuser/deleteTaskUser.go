package taskuser

import (
	"net/http"

	"github.com/RodrigoSousa101/ai_workflow/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RemoveUserFromTask(c *gin.Context) {
	taskID := c.Param("task_id")
	userID := c.Param("user_id")
	var taskUser models.TaskUser
	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("task_id = ? AND user_id = ?", taskID, userID).First(&taskUser).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "TaskUser not found"})
		return
	}

	if err := db.Delete(&taskUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to remover user from Task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User removed from task successfully"})
}
