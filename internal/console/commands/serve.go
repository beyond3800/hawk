package commands

import (
	"fmt"

	"github.com/beyond3800/hawk/config"
	"github.com/beyond3800/hawk/core/database"
	"github.com/beyond3800/hawk/routes"
	"github.com/spf13/cobra"
)

var serve = &cobra.Command{
	Use: "serve",
	Short: "Start the server at default \n",
	Run: func(cmd *cobra.Command, args []string) {
		
		database.ConnectDatabase()
		config.ConnectRedis()
		r := routes.SetupRoutes()
		r.Run(":8080")
		fmt.Printf("Running on server %v","127.0.0.1:8080")
	},
}



func ServeCommand() *cobra.Command {
	return serve
}