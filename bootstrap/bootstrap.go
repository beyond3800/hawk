package bootstrap

import (
	"os"
	"fmt"

	bootstrap "github.com/beyond3800/hawk/internal/bootstrap"
	"github.com/beyond3800/hawk/auth"
)

func Run() {
	if err := auth.Configure(auth.Config{
        SecretKey: os.Getenv("APP_SECRET_KEY"),
        Issuer:    os.Getenv("APP_ISSUER"),
    }); err != nil {
        fmt.Println(err)
    }
	bootstrap.Bootstrap()
}