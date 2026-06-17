package commands

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/beyond3800/hawk/lib"
	"github.com/spf13/cobra"
)

var makeModelCmd = &cobra.Command{
	Use:   "model",
	Short: "Create a new model",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		name := args[0]
		dir := "models"
		if !strings.HasSuffix(name,"Model"){
			name = name + "Model"
		}
		if lib.FileExist(dir,name){
			fmt.Println("This file already in" + dir)
			return
		}
		if err := os.MkdirAll(dir, 0755); err != nil {
			log.Fatal(err)
		}
		if err := lib.GenerateTemplate(name,"model"); err != nil{
			log.Println(err)
			return
		}

		log.Println("✅ Model created:")
	},
}


func ModelCommand() *cobra.Command {
	return makeModelCmd
}
