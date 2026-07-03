package commands

import (
	"github.com/beyond3800/hawk/internal/bootstrap"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show migration status",
	RunE: func(cmd *cobra.Command, args []string) error{
		return bootstrap.RunApp("hawk_status")
	},
}



func StatusCommand() *cobra.Command {
	return statusCmd
}