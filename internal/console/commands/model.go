package commands

import (
	"fmt"
	"log"
	"os"

	"github.com/beyond3800/hawk/lib"
	"github.com/spf13/cobra"
)


func createModel(name string, dir string){
	if lib.FileExist(dir,name){
		fmt.Println("This file already in" + dir)
		return
	}
	if err := os.MkdirAll(dir, 0755); err != nil {
		log.Fatal(err)
	}
	path := "app/Models"
	if err := lib.GenerateTemplate(name,"model",path); err != nil{
		log.Println(err)
		return
	}

	log.Println("✅ Model created:")
}

var makeModelCmd = &cobra.Command{
	Use:   "model",
	Short: "Create a new model",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		name := args[0]
		dir := "app/Models"
		createModel(name,dir)
	},
}


func ModelCommand() *cobra.Command {
	return makeModelCmd
}
