package Make

import (
	"fmt"
	"os"

	templatecreator "github.com/beyond3800/hawk/internal/templateCreator"
	"github.com/beyond3800/hawk/internal/lib"
	"github.com/spf13/cobra"
)

var isSeeder bool
var factoryCmd = &cobra.Command{
	Use:   "factory",
	Short: "Create a new factory",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dir := "database/factory"
		name := args[0]
		factoryName := name
		_,err := os.Stat(dir)
		if err !=nil{
			if os.IsNotExist(err) {
				if err := os.MkdirAll(dir, 0755); err != nil {
					fmt.Println("Unable to create resource directory:", err)
					return
				}
			}
		}
		if lib.FileExist(dir,name){
			fmt.Println(`This file already exist in ` + dir)
			return
		}
		if err := templatecreator.Factory(name, "factory", factoryName, dir); err != nil {
			fmt.Println(err)
		}
		fmt.Println("Factory created successfully")
		if isSeeder{
			CreateSeeder(name, "database/seeders")
		}
		
	},
}


func FactoryCommand() *cobra.Command {
	factoryCmd.Flags().BoolVarP(&isSeeder,"seeder","s",false,"Creating seeder controller")
	return factoryCmd
}