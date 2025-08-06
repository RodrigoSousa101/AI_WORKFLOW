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
	var workflowCreated []models.Workflow

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("user_id = ?", userID).Find(&workflowCreated).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "failed to retrieve workflows"})
		return
	}

	if len(workflowCreated) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "no workflows found for this user"})
		return
	}

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

	var WorkFlowsworker []models.Workflow

	if len(WorkFlowsID) > 0 {
		if err := db.Where("id IN ?", WorkFlowsID).Find(&WorkFlowsworker).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "failed to retrieve workflows"})
		}

	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "no WorkflowID found"})
		return
	}

	if len(WorkFlowsworker) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": " Workflows not found"})
	}

	c.JSON(http.StatusOK, gin.H{"Workflowsworker": WorkFlowsworker, "WorkflowCreated": workflowCreated})

}
