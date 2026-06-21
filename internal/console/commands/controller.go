package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/beyond3800/hawk/lib"
	"github.com/spf13/cobra"
)

// var isApi bool
var controller = &cobra.Command{
	Use: "controller",
	Short: "creates a controller in the controller's folder",
	Long: "",
	Run: func(cmd *cobra.Command, args []string) {
		dir := "app/Http/Controllers"
		_,err:= os.Stat(dir);
		if  os.IsNotExist(err) {
			err := os.Mkdir(dir,0755)
			if err !=nil {
				fmt.Println("Unable to create contoller folder")
			}
		}
		
		
		// var content string
		// if isApi{
			// fmt.Println("you are creating an api controller")
			// string(s[0])
			name:=args[0]
			if !strings.HasSuffix(name,"Controller") && !strings.HasSuffix(name,"controller"){
				name = name + "Controller"
			}
			if lib.FileExist(dir,name){
				fmt.Println(`This file already exist in ` + dir)
				return
			}

			err =lib.GenerateTemplate(name,"controllers",dir)
			if err != nil{
				fmt.Println(err)
				return
			}
			fmt.Println("Controller created successfully")
		// }
		// else{
			// txt,err:= os.ReadFile("./files/api.txt")
			// if err != nil{
			// 	fmt.Println("unable to read")
			// }
			// content = string(txt)
		// }
		// err = os.WriteFile("./controllers/"+args[0]+".go",[]byte(content),0644)
		// if err != nil {
		// 	fmt.Println("Unable to create controller")
		// }
		// fmt.Print("creating a controller........\n controller created successfully")
	},
}
func ControllerCommand() *cobra.Command {
	return controller
}