package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var migrationCmd = &cobra.Command{
	Use: "migration",
	Short: `
	      migration status
	      migration rollback
	      migration fresh
	      migration refresh
	      migration reset`,
	Long: "This command is used to run and manage database migrations.",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0  {
			fmt.Printf(`This are the commands for make
				migration status
				migration rollback`)
		}
		for _,flag:= range args{
			if flag == "migration"{
				fmt.Println("make migration")
			}
		}
	},
}

func MigrationCommands() *cobra.Command {
	return migrationCmd
}