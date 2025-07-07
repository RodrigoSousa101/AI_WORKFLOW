package workflow

import (
	"net/http"
	"time"

	"github.com/RodrigoSousa101/ai_workflow/models"
	"github.com/RodrigoSousa101/ai_workflow/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateWorkflow(c *gin.Context) {

	User, err := utils.GetUser(c)
	if err != nil {
		c.JSON(400, gin.H{"error": "User not found"})
	}

	var workflow models.Workflow
	if err := c.ShouldBindJSON(&workflow); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.MustGet("db").(*gorm.DB)
	workflow.ID = uuid.New()
	workflow.UserID = User.ID
	workflow.CreatedAt = time.Now()
	workflow.UpdatedAt = time.Now()

	if err := db.Create(&workflow).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Workflow"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "workflow created successfully", "Workflow": workflow})
}
