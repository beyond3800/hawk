package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func ConnectDatabase() error {
	var err error
	if err = godotenv.Load(); err != nil {
		return fmt.Errorf("Error loading .env file: %v", err)
	}
	
	driver := os.Getenv("DB_DRIVER")
	switch driver {
	case "mysql":
		if err := connectMySQL(Config{
				os.Getenv("DB_HOST"),
				os.Getenv("DB_PORT"),
				os.Getenv("DB_USER"),
				os.Getenv("DB_PASS"),
				os.Getenv("DB_NAME"),
			});
		err != nil{
			fmt.Println(err)
		}
		fmt.Println("db connected")
	}
	return nil
}