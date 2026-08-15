package Make

import (
	"fmt"
	"os"
	"strings"

	"github.com/beyond3800/hawk/internal/lib"
	templatecreator "github.com/beyond3800/hawk/internal/templateCreator"
	"github.com/spf13/cobra"
)

func CreateRequest(dir, name string) error {
	newName := strings.ToLower(name)
	
	if !strings.HasSuffix(newName, "request"){
		capName := templatecreator.ToTitle(name)
		newName = fmt.Sprintf("%vRequest", capName)
	}

	if lib.FileExist(dir, newName){
		return fmt.Errorf("This file exist in %v", dir)
	}
	err := templatecreator.MakeTemplate(name+".go", "request", dir, newName, "")
	if err != nil {
		return err
	}
	fmt.Println("Request created successfully")
	return nil
}

var requestCmd = &cobra.Command{
	Use:   "request",
	Short: "Create a new request",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dir := "app/Requests"
		name := args[0]
		requestName := name
		_,err := os.Stat(dir)
		if err !=nil{
			if os.IsNotExist(err) {
				if err := os.MkdirAll(dir, 0755); err != nil {
					fmt.Println("Unable to create request directory:", err)
					return
				}
			}
		}
		CreateRequest(dir, requestName)
	},
}

func RequestCommand() *cobra.Command {
	return requestCmd
}