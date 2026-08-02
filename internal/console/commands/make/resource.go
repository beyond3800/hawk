package Make

import (
	"fmt"
	"os"

	templatecreator "github.com/beyond3800/hawk/internal/templateCreator"
	"github.com/beyond3800/hawk/internal/lib"
	"github.com/spf13/cobra"
)


var resourceCmd = &cobra.Command{
	Use:   "resource",
	Short: "Create a new resource",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dir := "app/Resources"
		name := args[0]
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
		if err := templatecreator.MakeResourceTemplate(name, "resource", dir); err != nil {
			fmt.Println(err)
		}
		fmt.Println("Resource created successfully")
	},
}


func ResourceCommand() *cobra.Command {
	return resourceCmd
}