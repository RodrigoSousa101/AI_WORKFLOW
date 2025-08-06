package workflowuser

import (
	"net/http"

	"github.com/RodrigoSousa101/ai_workflow/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetWorkflowByUsers(c *gin.Context) {

	userID := c.Param("user_id")
	var workflowUser []models.WorkflowUser
	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("user_id = ?", userID).Find(&workflowUser).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "failed to retrieve workflowusers"})
		return
	}

	if len(workflowUser) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "no workflowusers found for this workflow"})
		return
	}

	var WorkFlowsID []uuid.UUID

	for _, wu := range workflowUser {
		WorkFlowsID = append(WorkFlowsID, wu.WorkflowID)
	}

	var WorkFlows []models.Workflow

	if len(WorkFlowsID) > 0 {
		if err := db.Where("id IN ?", WorkFlowsID).Find(&WorkFlows).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "failed to retrieve workflows"})
		}

	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "no WorkflowID found"})
		return
	}

	if len(WorkFlows) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": " Workflows not found"})
	}

	c.JSON(http.StatusOK, gin.H{"Workflows": WorkFlows})

}
