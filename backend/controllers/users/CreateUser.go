package controllers

import (
	"AI_WORKFLOW/backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateUser(c *gin.Context) {
	var User models.User
	if err := c.ShouldBindJSON(&User); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	User.ID = uuid.New().String()
	User.CreatedAt = models.GetCurrentTime()
	User.UpdatedAt = models.GetCurrentTime()

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Create(&User).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": User})
}
