package task

import (
	"net/http"

	"github.com/RodrigoSousa101/ai_workflow/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetTasksByWorkflow(c *gin.Context) {
	var tasks []models.Task
	var workflow models.Workflow
	WorkflowID := c.Param("workflow_id")
	db := c.MustGet("db").(*gorm.DB)

	if err := db.First(&workflow, "id = ?", WorkflowID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error: ": "workflow not found"})
		return
	}

	if err := db.Where("workflow_id = ?", WorkflowID).Find(&tasks).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "no tasks found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "tasks founded", "Tasks": tasks})
}
