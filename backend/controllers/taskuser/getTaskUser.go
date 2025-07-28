package taskuser

import (
	"net/http"

	"github.com/RodrigoSousa101/ai_workflow/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetTaskUser(c *gin.Context) {
	var taskUser models.TaskUser
	userID := c.Param("user_id")
	taskID := c.Param("task_id")
	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("task_id = ? AND user_id = ?", taskID, userID).First(&taskUser).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "taskuser not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "taskuser found", "worflowuser": taskUser})
}
