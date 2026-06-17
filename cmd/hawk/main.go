package main

import (
	"fmt"
	"os"

	"github.com/beyond3800/hawk/commands"
	"github.com/beyond3800/hawk/config"
	"github.com/beyond3800/hawk/core/database"
)

func main() {
	
	database.ConnectDatabase()
	config.ConnectRedis()
	
	if err := commands.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}