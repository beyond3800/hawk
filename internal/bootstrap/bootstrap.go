package bootstrap

import (
	"fmt"
	"strings"

	"github.com/beyond3800/hawk/auth"
	"github.com/beyond3800/hawk/internal/bootstrap/database"
	"github.com/beyond3800/hawk/internal/env"
	"github.com/beyond3800/hawk/storage"
)

func Bootstrap() {
    
    err := env.Load(".env")
    if err != nil{
        fmt.Println("Unable to load env")
    }
    secretKey, ok := env.Get("APP_SECRET_KEY")
    issuer, _ := env.Get("APP_ISSUER")
    if ok{
        auth.Configure(auth.Config{
            SecretKey: secretKey,
            Issuer: issuer,
        })
    }
    
    dbEnable,_ := env.Get("DB_ENABLED")
    redisEnabled,_ := env.Get("REDIS_ENABLED")
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

    	storageEnabled, _ := env.Get("STORAGE_ENABLED")
	if strings.EqualFold(storageEnabled, "true") {

		driver, _ := env.Get("STORAGE_DRIVER")
		root, _ := env.Get("STORAGE_ROOT")
		baseURL, _ := env.Get("STORAGE_URL")

		err := storage.Configure(storage.Config{
			Driver:  driver,
			Root:    root,
			BaseURL: baseURL,
		})

		if err != nil {
			fmt.Println("Storage:", err)
		}
	}
}