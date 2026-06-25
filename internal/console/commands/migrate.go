package commands

import (
	"fmt"

	_ "github.com/beyond3800/hawk/database/migrations"
	"github.com/beyond3800/hawk/internal/boostrap/database"
	"github.com/beyond3800/hawk/internal/console/migration"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var migrateCmd  = &cobra.Command{
	Use: "migrate",
	Short: "Migrate all table to database",
	Run: func(cmd *cobra.Command, args []string) {
		if err := godotenv.Load(); err != nil{
			fmt.Println("Error loading .env file")
			return
		}
		
		database.ConnectDatabase()
		if err := migration.Run(); err != nil{
			fmt.Println(err)
			fmt.Println("Unable to migrate database")
			return
		}
	},
}


func MigrateCommand() *cobra.Command {
	return migrateCmd
}