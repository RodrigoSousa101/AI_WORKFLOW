package users

import (
	"github.com/gin-gonic/gin"
)

func UserRoutes(rg *gin.RouterGroup) {
	router := rg.Group("/users")
	{
		router.POST("", CreateUser)
	}
}
