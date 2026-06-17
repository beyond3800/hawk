package util

import (
	"os"

	"github.com/joho/godotenv"
)

func UserCookieName() (string,error){
	err :=godotenv.Load()
	if err != nil{
		return "",err
	}
	return os.Getenv("USER_ID"), nil
}