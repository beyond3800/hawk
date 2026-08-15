package Make

import (
	"fmt"
	"os"

	"github.com/beyond3800/hawk/internal/lib"
	"github.com/spf13/cobra"
)

func CreateMigration(name string, migrationDir string) {
	_,err := os.Stat(migrationDir)
		if err !=nil{
			if os.IsNotExist(err) {
				if err := os.MkdirAll(migrationDir, 0755); err != nil {
					fmt.Println("Unable to create migration directory:", err)
					return
				}
			}
		}
		
		if name == "" {
			fmt.Println("Migration name is required")
			return
		}
		if lib.FileExist(migrationDir,name){
			fmt.Println("This file already in" + migrationDir)
			return
		}
		
		err = lib.MakeMigrationTemplate(name,"migration",name, migrationDir)
		if err != nil{
			fmt.Println("Unable to create file")
			return
		}
		fmt.Println("Migration created successfully")
}

var migrationCmd = &cobra.Command{
	Use: "migration",
	Short: "Make migration table",
	Long: "This command is use to make migration file in the database/migrations folder\nUsage: make migration [name]",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Creating migration...")
		migrationDir := "database/migrations"
		name := args[0]

		CreateMigration(name,migrationDir)
	},
}


func MigrationCommand() *cobra.Command {
	return migrationCmd
}
