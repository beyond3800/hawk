package database

import (
	"fmt"
	"os"

)

func ConnectDatabase() error {
	
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
	}
	return nil
}