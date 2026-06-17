package commands

import (
	"fmt"

	"github.com/beyond3800/hawk/core/migration"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show migration status",
	Run: func(cmd *cobra.Command, args []string) {

		if err := migration.Status(); err != nil {
			fmt.Println(err)
		}
	},
}



func StatusCommand() *cobra.Command {
	return statusCmd
}