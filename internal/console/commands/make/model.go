package Make

import (
	"fmt"
	"log"
	"os"

	"github.com/beyond3800/hawk/internal/lib"
	"github.com/spf13/cobra"
)

var isMigration bool

func CreateModel(name string, dir string){
	fmt.Println("Createing model ....")
	if lib.FileExist(dir,name){
		fmt.Println("This file already in" + dir)
		return
	}
	if err := os.MkdirAll(dir, 0755); err != nil {
		log.Fatal(err)
	}
	if err := lib.GenerateTemplate(name,"model",dir); err != nil{
		log.Println(err)
		return
	}

	log.Println("✅ Model created:")
}

var modelCmd = &cobra.Command{
	Use:   "model [name]",
	Short: "Create a new model",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		name := args[0]
		dir := "app/Models"
		CreateModel(name,dir)
		if isMigration{
			CreateMigration(name, "database/migrations")
		}
	},
}


func ModelCommand() *cobra.Command {
	modelCmd.Flags().BoolVarP(&isMigration,"migration","m",false,"Creating migration table")
	return modelCmd
}
