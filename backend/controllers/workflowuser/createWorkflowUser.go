package workflowuser

import (
	"net/http"

	"github.com/RodrigoSousa101/ai_workflow/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateWorkflowUser(c *gin.Context) {
	var workflowUser models.WorkflowUser
	db := c.MustGet("db").(*gorm.DB)

	if err := c.ShouldBindJSON(&workflowUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var workflow models.Workflow
	if err := db.Where("id = ?", workflowUser.WorkflowID).First(&workflow).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Workflow not found"})
	}

	var user models.User
	if err := db.Where("id = ?", workflowUser.UserID).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
	}
}
