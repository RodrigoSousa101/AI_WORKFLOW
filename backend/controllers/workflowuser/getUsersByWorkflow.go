package workflowuser

import (
	"net/http"

	"github.com/RodrigoSousa101/ai_workflow/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetUsersByWorkflow(c *gin.Context) {

	workflowID := c.Param("workflow_id")
	var workflowUser []models.WorkflowUser
	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("workflow_id = ?", workflowID).Find(&workflowUser).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "failed to retrieve workflowusers"})
		return
	}

	if len(workflowUser) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "no workflowusers found for this workflow"})
		return
	}

	var userIDs []uuid.UUID

	for _, wu := range workflowUser {
		userIDs = append(userIDs, wu.UserID)
	}

	var users []models.User

	if len(userIDs) > 0 {
		if err := db.Where("id IN ?", userIDs).Find(&users).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "failed to retrieve users"})
		}

	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "no usersID found"})
		return
	}

	if len(users) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": " users not found"})
	}

	c.JSON(http.StatusOK, gin.H{"workers": users})

}
