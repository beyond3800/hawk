package bootstrap

import (
	"fmt"
	"os"

	"github.com/beyond3800/hawk/internal/boostrap/database"
)

func Boostrap() {
	dbEnable := os.Getenv("DB_ENABLED")
	redisEnable := os.Getenv("REDIS_ENABLED")

	if dbEnable != "" {
		err := database.ConnectDatabase()
		fmt.Println(err)
	}
	if redisEnable != ""{
		err := ConnectRedis()
		fmt.Println(err)
	}
}