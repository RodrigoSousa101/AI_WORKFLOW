package workflow

import (
	"github.com/gin-gonic/gin"
)

func UserRoutes(rg *gin.RouterGroup) {
	router := rg.Group("/workflow")
	{
		router.POST("", CreateWorkflow)
		router.GET("", GetAllWorkflows)
		router.GET("/:id", GetWorkflow)
		router.PUT("/:id", UpdateWorkflow)
		router.DELETE("/:id", DeleteWorkflow)
		router.POST("/:id", AddWorkerToWorkflow)
	}
}
