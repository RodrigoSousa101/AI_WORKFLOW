package users

import (
	"github.com/gin-gonic/gin"
)

func UserRoutes(rg *gin.RouterGroup) {
	router := rg.Group("/users")
	{
		router.POST("", CreateUser)
		router.GET("/:id", GetUser)
		router.GET("", GetAllUsers)
		router.PUT("/:id", UpdateUser)
		router.DELETE("/:id", DeleteUser)
		router.GET("/current", GetCurrentUser)
	}
}
