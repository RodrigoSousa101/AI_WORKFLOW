package workflow

import (
	"net/http"

	"github.com/RodrigoSousa101/ai_workflow/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddWorkerToWorkflow(c *gin.Context) {
	WorkflowID := c.Param("id")
	var workflow models.Workflow
	var worker models.User
	db := c.MustGet("db").(*gorm.DB)

	if err := c.ShouldBindJSON(&worker); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.First(&workflow, "id = ?", WorkflowID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "workflow not found"})
		return
	}

	if err := db.Where("email = ?", worker.Email).First(&worker).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "no worker founded"})
		return
	}

	if err := db.Model(&workflow).Association("Workers").Append(&worker); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to associate user to workflow"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user added to workflow", "Workflow": workflow})
}
