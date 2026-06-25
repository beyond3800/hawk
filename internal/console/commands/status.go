package commands

import (
	"fmt"

	"github.com/beyond3800/hawk/internal/boostrap/database"
	"github.com/beyond3800/hawk/internal/console/migration"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show migration status",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		err = database.ConnectDatabase()
		err = migration.Status()
		if err != nil {
			fmt.Println(err)
		}
	},
}



func StatusCommand() *cobra.Command {
	return statusCmd
}