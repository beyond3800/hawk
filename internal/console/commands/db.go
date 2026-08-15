package commands



import (
	"fmt"

	"github.com/spf13/cobra"
)

var dbCmd = &cobra.Command{
	Use: "db",
	Short: `
	      db seed`,
	Long: "This command is use for database services to seed the database",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0  {
			fmt.Printf(`This are the commands for make db seed`)
		}

	},
}

func DbCommand() *cobra.Command {
	return dbCmd
}