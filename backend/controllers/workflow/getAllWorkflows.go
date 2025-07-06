package workflow

import (
	"net/http"

	"github.com/RodrigoSousa101/ai_workflow/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllWorkflows(c *gin.Context) {
	var workflow []models.Workflow
	db := c.MustGet("db").(*gorm.DB)

	if err := db.Find(&workflow).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve workflow"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": workflow})

}
