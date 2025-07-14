package workflow

import (
	"net/http"
	"time"

	"github.com/RodrigoSousa101/ai_workflow/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateWorkflow(c *gin.Context) {
	var workflow models.Workflow
	var existingWorkflow models.Workflow
	WorkflowID := c.Param("id")
	if err := c.ShouldBindJSON(&workflow); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("id = ?", WorkflowID).First(&existingWorkflow).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Workflow not found"})
		return
	}

	if err := db.Model(&existingWorkflow).Where("id = ?", WorkflowID).Updates(workflow).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update Workflow"})
		return
	}
	existingWorkflow.UpdatedAt = time.Now()
	c.JSON(http.StatusOK, gin.H{"message": "Workflow updated successfully", "Workflow": existingWorkflow})

}
