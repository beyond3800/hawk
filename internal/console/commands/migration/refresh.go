package migration

import (
	"github.com/beyond3800/hawk/internal/bootstrap"
	"github.com/spf13/cobra"
)

var refresh = &cobra.Command{
	Use:   "refresh",
	Short: "Refresh migration",
	RunE: func(cmd *cobra.Command, args []string) error{
		return bootstrap.RunApp("hawk_refresh")
	},
}

func ReFreshCommand() *cobra.Command {
	return refresh
}