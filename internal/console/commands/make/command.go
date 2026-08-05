package Make

import (
	"fmt"
	"os"
	"strings"

	"github.com/beyond3800/hawk/internal/lib"
	templatecreator "github.com/beyond3800/hawk/internal/templateCreator"
	"github.com/spf13/cobra"
)

var commandCmd = &cobra.Command{
	Use:   "command",
	Short: "Create a new command",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dir := "internal/console"
		name := args[0]
		_,err := os.Stat(dir)
		if err !=nil{
			if os.IsNotExist(err) {
				if err := os.MkdirAll(dir, 0755); err != nil {
					fmt.Println("Unable to create command:", err)
					return
				}
			}
		}
		if lib.FileExist(dir,name){
			fmt.Println(`This file already exist in ` + dir)
			return
		}
		name = strings.ToLower(name)
		if err := templatecreator.MakeTemplate(name+".go", "command", dir, templatecreator.ToTitle(name), "" ); err != nil {
			fmt.Println(err)
		}
		fmt.Println("Command created successfully")
	},
}


func Command() *cobra.Command {
	return commandCmd
}