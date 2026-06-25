package commands

import (
	"fmt"

	_ "github.com/beyond3800/hawk/database/migrations"
	"github.com/beyond3800/hawk/internal/boostrap/database"
	"github.com/beyond3800/hawk/internal/console/migration"

	"github.com/spf13/cobra"
)

var rollback = &cobra.Command{
	Use:   "rollback",
	Short: "Rollback latest migration batch",
	Run: func(cmd *cobra.Command, args []string) {
		database.ConnectDatabase()
		if err := migration.Rollback(); err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Rollback complete")
	},
}




func RollbackCommand() *cobra.Command {
	return rollback
}