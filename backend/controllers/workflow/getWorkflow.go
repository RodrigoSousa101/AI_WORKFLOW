package workflow

import (
	"net/http"

	"github.com/RodrigoSousa101/ai_workflow/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetWorkflow(c *gin.Context) {
	WorkflowID := c.Param("id")
	var workflow models.Workflow
	db := c.MustGet("db").(*gorm.DB)

	if err := db.Preload("Workers").First(&workflow, "id = ?", WorkflowID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "workflow not found"})
		return
	}

	c.JSON(200, gin.H{"message": "Wokflow found", "Workflow": workflow})
}
