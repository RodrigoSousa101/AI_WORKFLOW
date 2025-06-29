package controllers

import (
	"gin-gonic/gin"
)

func UserRoutes(c *gin.Context) {
	router := c.Group("/users")
	{
		router.POST("", CreateUser)
	}
}
