package database

import (
	"fmt"

	"github.com/beyond3800/hawk/internal/env"
)

func ConnectDatabase() error {
	
	host,_ := env.Get("DB_HOST")
	port,_ := env.Get("DB_PORT")
	user,_ := env.Get("DB_USER")
	pwd,_ := env.Get("DB_PASS")
	name,_ := env.Get("DB_NAME")
	driver,_ := env.Get("DB_DRIVER")

	switch driver {
	case "mysql":
		if err := connectMySQL(Config{
				host,
				port,
				user,
				pwd,
				name,
			});
		err != nil{
			fmt.Println(err)
		}
	}
	return nil
}