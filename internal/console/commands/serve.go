package commands

import (
	"os"
	"strings"

	bootstrap "github.com/beyond3800/hawk/internal/bootstrap"
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
		os.Setenv("APP_PORT", port)
		bootstrap.Air()
	},
}



func ServeCommand() *cobra.Command {
	return serve
}