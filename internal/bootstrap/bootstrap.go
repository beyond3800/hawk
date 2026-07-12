package bootstrap

import (
	"fmt"
	"os"
	"strings"

	"github.com/beyond3800/hawk/auth"
	"github.com/beyond3800/hawk/internal/bootstrap/database"
	"github.com/joho/godotenv"
)

func Bootstrap() {
    if err := godotenv.Load(); err != nil {
        fmt.Println("Warning:", err)
    }
    if err := auth.Configure(auth.Config{
        SecretKey: os.Getenv("APP_SECRET_KEY"),
        Issuer:    os.Getenv("APP_ISSUER"),
    }); err != nil {
        fmt.Println("Auth:", err)
    }
    dbEnable := os.Getenv("DB_ENABLED")
    redisEnabled := os.Getenv("REDIS_ENABLED")
    if strings.ToLower(dbEnable) == "true" {
        if err := database.ConnectDatabase(); err != nil {
            fmt.Println("Database:", err)
        }
    }
    if strings.ToLower(redisEnabled) == "true" {
        if err := ConnectRedis(); err != nil {
            fmt.Println("Redis:", err)
        }
    }
}