package commands

import (

	"github.com/beyond3800/hawk/internal/bootstrap"
	"github.com/spf13/cobra"
)



var migrateCmd  = &cobra.Command{
	Use: "migrate",
	Short: "Migrate all table to database",
	RunE: func(cmd *cobra.Command, args []string) error{
		return bootstrap.RunApp("hawk_migrate")
	},
}


func MigrateCommand() *cobra.Command {
	return migrateCmd
}