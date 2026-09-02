package migration

import (
	"github.com/beyond3800/hawk/internal/bootstrap"
	"github.com/spf13/cobra"
)

var rollback = &cobra.Command{
	Use:   "rollback",
	Short: "Rollback latest migration batch",
	RunE: func(cmd *cobra.Command, args []string) error{
		return bootstrap.RunApp("hawk_rollback")
	},
}

func RollbackCommand() *cobra.Command {
	return rollback
}