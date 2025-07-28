package workflowuser

import (
	"net/http"

	"github.com/RodrigoSousa101/ai_workflow/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllWorkflowUsers(c *gin.Context) {
	var workflowuser []models.WorkflowUser
	db := c.MustGet("db").(*gorm.DB)

	if err := db.Find(&workflowuser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve workflowuser"})
		return
	}

	if len(workflowuser) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No workflowusers found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"workflowusers": workflowuser})

}
