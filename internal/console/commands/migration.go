package commands

import (
	"fmt"
	"os"
	"time"

	"github.com/beyond3800/hawk/lib"
	"github.com/spf13/cobra"
)

func createMigration(name string, migrationDir string){
	_,err := os.Stat(migrationDir)
		if err !=nil{
			if err := os.Mkdir(migrationDir, 0755); err != nil{
				fmt.Println("Unable to create file")
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
		timestamp := time.Now().Format("20060102150405")
		migrationName := fmt.Sprintf("%s_%s", timestamp, name)
		err = lib.MakeMigrationTemplate(name,"migration",migrationName)
		if err != nil{
			fmt.Println("Unable to create file")
			return
		}
}

var migrationCmd = &cobra.Command{
	Use: "migration",
	Short: "Make migration table",
	Long: "This command is use to make migration file in the database/migrations folder\nUsage: make migration [name]",
	// Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Creating migration...")
		migrationDir := "database/migrations"
		name := args[0]

		createMigration(name,migrationDir)
	},
}


func MigrationCommand() *cobra.Command {
	return migrationCmd
}
