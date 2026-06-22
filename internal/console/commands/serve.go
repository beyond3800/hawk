package commands

import (
	"github.com/beyond3800/hawk"
	"github.com/spf13/cobra"
)

var serve = &cobra.Command{
	Use:   "serve",
	Short: "Start server",

	Run: func(cmd *cobra.Command, args []string) {
		port := "8080"

		if len(args) > 0 {
			port = args[0]
		}

		app := hawk.New()
		app.Run(":" + port)
	},
}



func ServeCommand() *cobra.Command {
	return serve
}