package main

import (
	"github.com/RodrigoSousa101/ai_workflow/controllers/auth"
	"github.com/RodrigoSousa101/ai_workflow/controllers/users"
	"github.com/RodrigoSousa101/ai_workflow/models"

	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

func main() {
	// Carrega .env (se existir)
	err := godotenv.Load()
	if err != nil {
		log.Println("Aviso: .env não carregado (provavelmente dentro do Docker)")
	}

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("Variável DATABASE_URL não configurada")
	}

	var db *gorm.DB
	maxAttempts := 10
	for i := 1; i <= maxAttempts; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		log.Printf("Tentativa %d/%d: erro ao conectar ao banco, tentando novamente em 2s...", i, maxAttempts)
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados após várias tentativas:", err)
	}

	log.Println("Conexão com o banco de dados bem-sucedida!")
	db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)

	// Migrar várias tabelas de uma vez:
	if err := db.AutoMigrate(&models.User{}, &models.Workflow{}, &models.Task{}); err != nil {
		log.Fatal("Erro ao migrar tabelas:", err)
	}

	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	workflowGroup := r.Group("/workflow")

	workflowGroup.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Siga"})
	})

	// Agora as rotas de usuário ficam em /workflow/users
	users.UserRoutes(workflowGroup)
	auth.AuthRoutes(workflowGroup)

	err = r.Run(":8080")
	if err != nil {
		log.Fatal("Erro ao iniciar o servidor:", err)
	}
}
