package workflowuser

import (
	"github.com/gin-gonic/gin"
)

func WorkflowUserRoutes(rg *gin.RouterGroup) {
	router := rg.Group("/workflowuser")
	{
		router.POST("", CreateWorkflowUser)
		router.GET("", GetAllWorkflowUsers)
		router.GET("/:workflow_id/:user_id", GetWorkflowUser)
		router.GET("/:workflow_id", GetUsersByWorkflow)
		router.DELETE("/:workflow_id/:user_id", RemoveUserFromWorkflow)
	}
}
