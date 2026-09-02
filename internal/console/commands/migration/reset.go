package migration

import (
	"github.com/beyond3800/hawk/internal/bootstrap"
	"github.com/spf13/cobra"
)

var reset = &cobra.Command{
	Use:   "reset",
	Short: "Reset migration",
	RunE: func(cmd *cobra.Command, args []string) error{
		return bootstrap.RunApp("hawk_reset")
	},
}

func ResetCommand() *cobra.Command {
	return reset
}