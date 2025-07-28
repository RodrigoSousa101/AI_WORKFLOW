package workflowuser

import (
	"net/http"

	"github.com/RodrigoSousa101/ai_workflow/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RemoveUserFromWorkflow(c *gin.Context) {
	WorkflowID := c.Param("workflow_id")
	UserID := c.Param("user_id")
	var workflowUser models.WorkflowUser
	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("workflow_id = ? AND user_id = ?", WorkflowID, UserID).First(&workflowUser).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "WorkflowUser not found"})
		return
	}

	if err := db.Delete(&workflowUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to remover user from workflow"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User removed from workflow successfully"})
}
