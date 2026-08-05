package commands

import (
	"fmt"

	"github.com/beyond3800/hawk/internal/bootstrap"
	"github.com/spf13/cobra"
)



var appCmd = &cobra.Command{
	Use:   "app",
	Short: "Run application console commands",
	RunE: func(cmd *cobra.Command, args []string) error {

		if len(args) == 0 {
			return fmt.Errorf("no application command supplied")
		}

		return bootstrap.RunApp(args...)
	},
}

func AppCommand() *cobra.Command {
	return appCmd
}