package util

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetFromEnv(name string) (string, error) {
	if err := godotenv.Load(); err != nil{
		return "",err
	}
	isExist := os.Getenv(name)
	if isExist == "" {
		return "", fmt.Errorf("This cookie does not exist")
	}
	return isExist,nil

}