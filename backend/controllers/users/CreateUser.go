package users

import (
	"net/http"
	"time"

	"github.com/RodrigoSousa101/ai_workflow/models"
	"github.com/RodrigoSousa101/ai_workflow/utils"

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
	User.ID = uuid.New()
	User.CreatedAt = time.Now()
	User.UpdatedAt = time.Now()
	User.Password, _ = utils.HashPassword(User.Password)

	db := c.MustGet("db").(*gorm.DB)

	if db.First(&models.User{}, "email = ?", User.Email).Error == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
		return
	}

	if err := db.Create(&User).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	User.Password = ""

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": User})
}
