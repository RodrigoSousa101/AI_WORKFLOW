package main

import (
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

    r := gin.Default()

    r.GET("/workflow", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "Siga"})
    })

    err = r.Run(":8080")
    if err != nil {
        log.Fatal("Erro ao iniciar o servidor:", err)
    }
}
