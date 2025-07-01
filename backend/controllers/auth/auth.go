package auth

import (
	"github.com/gin-gonic/gin"
)

func AuthRoutes(rg *gin.RouterGroup) {
	router := rg.Group("/auth")
	{
		router.POST("/login", Login)
	}
}
