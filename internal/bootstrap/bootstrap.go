package bootstrap

import (
	"fmt"
	"os"
	"strings"

	"github.com/beyond3800/hawk/internal/bootstrap/database"
)

func Bootstrap() {
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