package workflow

import (
	"github.com/RodrigoSousa101/ai_workflow/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetWorkflow(c *gin.Context) {
	WorkflowID := c.Param("id")
	var workflow models.Workflow
	db := c.MustGet("db").(*gorm.DB)

	if err := db.First(&workflow, "id = ?", WorkflowID).Error; err != nil {
		c.JSON(404, gin.H{"error": "Workflow not found"})
		return
	}

	c.JSON(200, gin.H{"message": "Wokflow found", "Workflow": workflow})
}
