package commands

import (
	"fmt"
	"strings"

	bootstrap "github.com/beyond3800/hawk/internal/boostrap"
	"github.com/beyond3800/hawk/internal/boostrap/database"
	"github.com/beyond3800/hawk/routes"
	"github.com/spf13/cobra"
)

var serve = &cobra.Command{
	Use: "serve",
	Short: "Start the server at default \n",
	Run: func(cmd *cobra.Command, args []string) {
		port := ":8080"

		if len(args) > 0 {
			if !strings.HasPrefix(args[0], ":") {
				port = ":" + args[0]
			} else {
				port = args[0]
			}
		}
		database.ConnectDatabase()
		bootstrap.ConnectRedis()
		r := routes.SetupRoutes()
		r.Run(port)
		fmt.Printf("Running on server %v","127.0.0.1:8080")
	},
}



func ServeCommand() *cobra.Command {
	return serve
}