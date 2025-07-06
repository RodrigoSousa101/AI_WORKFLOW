package middleware

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/RodrigoSousa101/ai_workflow/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var tokenString string

		// 1. Tenta pegar o token do Header
		authHeader := c.GetHeader("Authorization")
		if strings.HasPrefix(authHeader, "Bearer ") {
			tokenString = strings.TrimPrefix(authHeader, "Bearer ")
		} else {
			// 2. Se não tiver no header, tenta pegar do cookie
			cookie, err := c.Cookie("Authorization")
			if err != nil {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			tokenString = cookie
		}

		// 3. Parse do token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY")), nil
		}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

		if err != nil || !token.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// 4. Validar claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// 5. Buscar usuário no banco
		db := c.MustGet("db").(*gorm.DB)
		var user models.User
		userID, ok := claims["sub"].(string)
		if !ok {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if err := db.Where("id = ?", userID).First(&user).Error; err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Adiciona o usuário no contexto
		c.Set("user", user)
		c.Next()
	}
}
