package task

import (
	"net/http"
	"time"

	"github.com/RodrigoSousa101/ai_workflow/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateTask(c *gin.Context) {
	var task models.Task
	var workflow models.Workflow
	WorkflowID := c.Param("workflow_id")
	db := c.MustGet("db").(*gorm.DB)

	parseUuid, err := uuid.Parse(WorkflowID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error converting string to uuid": err.Error()})
		return
	}

	if err := db.First(&workflow, "id = ?", WorkflowID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error:": "workflow not found"})
		return
	}

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task.ID = uuid.New()
	task.WorkflowID = parseUuid
	task.Created_at = time.Now()
	task.Updated_at = time.Now()

	if err := db.Create(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "error creating task "})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "task created sucefully", "task": task})

}
