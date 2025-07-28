package workflowuser

import (
	"net/http"

	"github.com/RodrigoSousa101/ai_workflow/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetWorkflowUser(c *gin.Context) {
	var workflowUser models.WorkflowUser
	userID := c.Param("user_id")
	workflowID := c.Param("workflow_id")
	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("workflow_id = ? AND user_id = ?", workflowID, userID).First(&workflowUser).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "workflowuser not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "workflowuser found", "worflowuser": workflowUser})
}
