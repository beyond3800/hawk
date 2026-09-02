package migration

import (
	"github.com/beyond3800/hawk/internal/bootstrap"
	"github.com/spf13/cobra"
)

var fresh = &cobra.Command{
	Use:   "fresh",
	Short: "Drops tables and refresh migration",
	RunE: func(cmd *cobra.Command, args []string) error{
		return bootstrap.RunApp("hawk_fresh")
	},
}

func FreshCommand() *cobra.Command {
	return fresh
}