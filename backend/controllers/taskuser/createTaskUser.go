package taskuser

import (
	"net/http"

	"github.com/RodrigoSousa101/ai_workflow/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateTaskUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var taskUser models.TaskUser

	type CreateInfo struct {
		UserID string    `json:"user_id"`
		TaskID uuid.UUID `json:"task_id"`
	}

	var input CreateInfo

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var task models.Task
	if err := db.Where("id = ?", input.TaskID).First(&task).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	taskUser.TaskID = input.TaskID

	var user models.User
	if err := db.Where("id = ?", input.UserID).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	taskUser.UserID = user.ID

	if err := db.Create(&taskUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create TaskUser"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "TaskUser created successfully", "WorkflowUser": taskUser})

}
