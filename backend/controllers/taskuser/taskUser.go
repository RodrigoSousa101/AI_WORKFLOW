package taskuser

import (
	"github.com/gin-gonic/gin"
)

func TaskUserRoutes(rg *gin.RouterGroup) {
	router := rg.Group("/taskuser")
	{
		router.POST("", CreateTaskUser)
		router.GET("", getAllTaskUser)
		router.GET("/:task_id/:user_id", GetTaskUser)
		//router.GET("/:workflow_id", GetUsersByWorkflow)
		router.DELETE("/:task_id/:user_id", RemoveUserFromTask)
	}
}
