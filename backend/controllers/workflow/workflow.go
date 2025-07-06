package workflow

import (
	"github.com/gin-gonic/gin"
)

func UserRoutes(rg *gin.RouterGroup) {
	router := rg.Group("/workflow")
	{
		router.POST("", CreateWorkflow)
		router.GET("", GetAllWorkflows)

	}
}
