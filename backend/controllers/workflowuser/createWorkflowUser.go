package workflowuser

import (
	"net/http"

	"github.com/RodrigoSousa101/ai_workflow/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateWorkflowUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var WorkflowUser models.WorkflowUser

	type CreateInfo struct {
		UserEmail  string    `json:"user_email"`
		WorkflowID uuid.UUID `json:"workflow_id"`
	}

	var input CreateInfo

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var workflow models.Workflow
	if err := db.Where("id = ?", input.WorkflowID).First(&workflow).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Workflow not found"})
		return
	}

	WorkflowUser.WorkflowID = input.WorkflowID

	var user models.User
	if err := db.Where("email = ?", input.UserEmail).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	WorkflowUser.UserID = user.ID

	if err := db.Create(&WorkflowUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create WorkflowUser"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "workflowUser created successfully", "WorkflowUser": WorkflowUser})

}
