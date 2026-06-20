package commands

import (
	"fmt"

	"github.com/beyond3800/hawk/core/migration"
	_ "github.com/beyond3800/hawk/database/migrations"
	bootstrap "github.com/beyond3800/hawk/internal/boostrap"
	"github.com/spf13/cobra"
)

var rollback = &cobra.Command{
	Use:   "rollback",
	Short: "Rollback latest migration batch",
	Run: func(cmd *cobra.Command, args []string) {
		bootstrap.DB()
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