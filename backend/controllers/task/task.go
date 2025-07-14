package task

import (
	"github.com/gin-gonic/gin"
)

func UserRoutes(rg *gin.RouterGroup) {
	router := rg.Group("/task")
	{
		router.POST("/:workflow_id", CreateTask)
		router.GET("", GetAllTasks)
		router.GET("/:id", GetTask)
		router.PUT("/:id", UpdateTask)
		router.DELETE("/:id", DeleteTask)
		router.GET("/tasksbyworkflow/:workflow_id", GetTasksByWorkflow)
	}
}
