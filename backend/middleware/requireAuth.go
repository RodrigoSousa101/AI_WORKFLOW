package middleware

import (
	"fmt"
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
		// 1. Pega o token do header Authorization
		authHeader := c.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// 2. Parse do token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("ACCESS_SECRET")), nil
		}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

		if err != nil || !token.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// 3. Validar claims e expiração
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		exp, ok := claims["exp"].(float64)
		if !ok || float64(time.Now().Unix()) > exp {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// 4. Buscar usuário no banco
		userID, ok := claims["sub"].(string)
		if !ok {
			// Se sub não é string, tenta converter de float64 para string (caso o ID seja numérico)
			if idFloat, ok := claims["sub"].(float64); ok {
				userID = fmt.Sprintf("%.0f", idFloat)
			} else {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
		}

		db := c.MustGet("db").(*gorm.DB)
		var user models.User
		if err := db.Where("id = ?", userID).First(&user).Error; err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// 5. Coloca o usuário no contexto para o handler usar
		c.Set("user", user)

		c.Next()
	}
}
