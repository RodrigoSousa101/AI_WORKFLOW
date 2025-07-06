package workflow

import (
	"github.com/RodrigoSousa101/ai_workflow/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteWorkflow(c *gin.Context) {
	var Workflow models.Workflow
	WorkflowID := c.Param("id")
	db := c.MustGet("db").(*gorm.DB)
	if db.First(&Workflow, "id = ?", WorkflowID).Error != nil {
		c.JSON(404, gin.H{"error": "Workflow not found"})
		return
	}
	if err := db.Delete(&Workflow, "id = ?", WorkflowID).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete Workflow"})
		return
	}

	c.JSON(200, gin.H{"message": "Workflow deleted successfully", "Workflow": Workflow})

}
