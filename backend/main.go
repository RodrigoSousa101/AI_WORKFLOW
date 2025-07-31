package main

import (
	"github.com/RodrigoSousa101/ai_workflow/controllers/auth"
	"github.com/RodrigoSousa101/ai_workflow/controllers/task"
	"github.com/RodrigoSousa101/ai_workflow/controllers/taskuser"
	"github.com/RodrigoSousa101/ai_workflow/controllers/users"
	"github.com/RodrigoSousa101/ai_workflow/controllers/workflow"
	"github.com/RodrigoSousa101/ai_workflow/controllers/workflowuser"
	"github.com/RodrigoSousa101/ai_workflow/middleware"
	"github.com/RodrigoSousa101/ai_workflow/models"

	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
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

	if err := db.AutoMigrate(&models.User{}, &models.Workflow{}, &models.Task{}, &models.WorkflowUser{}, &models.TaskUser{}); err != nil {
		log.Fatal("Erro ao migrar tabelas:", err)
	}

	r := gin.Default()
	r.Use(cors.Default())
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	workflowGroup := r.Group("/api")

	// Agora as rotas de usuário ficam em /workflow/users

	auth.AuthRoutes(workflowGroup)
	users.UserRoutes(workflowGroup)
	workflowGroup.Use(middleware.RequireAuth())
	{
		//users.UserRoutes(workflowGroup)
		workflow.WorkflowRoutes(workflowGroup)
		task.TaskRoutes(workflowGroup)
		workflowuser.WorkflowUserRoutes(workflowGroup)
		taskuser.TaskUserRoutes(workflowGroup)
	}

	err = r.Run(":8080")
	if err != nil {
		log.Fatal("Erro ao iniciar o servidor:", err)
	}
}
