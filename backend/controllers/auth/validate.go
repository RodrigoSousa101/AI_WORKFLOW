package auth

import (
	"github.com/gin-gonic/gin"
)

func Validate(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(200, gin.H{
		"message": user})
}
