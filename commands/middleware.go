package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/beyond3800/hawk/lib"
	"github.com/spf13/cobra"
)

var middlewareCmd = &cobra.Command{
	Use:   "middleware",
	Short: "Run middleware commands",
	Run: func(cmd *cobra.Command, args []string) {
		// This command is a placeholder for middleware-related commands.
		dir := "middleware"
		_,err:= os.Stat(dir);
		if  os.IsNotExist(err) {
			err := os.Mkdir(dir,0755)
			if err !=nil {
				fmt.Println("Unable to create middleware folder")
			}
		}
		
		name:=args[0]
		if !strings.HasSuffix(name,"Middleware"){
				name = name + "Middleware"
			}
		fmt.Println("Creating middleware:", name)
		if lib.FileExist(dir,name){
			fmt.Println("This file already in" + dir)
			return
		}
		err =lib.MakeMiddlewareTemplate(name,"middleware")
		if err != nil{
			fmt.Println(err)
			return
		}
	fmt.Println("Middleware created successfully", name)
	},
}


func MiddlewareCommand () *cobra.Command{
	return middlewareCmd
}


