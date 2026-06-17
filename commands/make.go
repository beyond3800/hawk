package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var make = &cobra.Command{
	Use: "make",
	Short: "make controller\nmake model [name] \nmake service [name] \nmake repository [name]\nmake all",
	Long: "This command is use to make model, controller, and so on",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0  {
			fmt.Printf(`This are the commands for make
				make:controller [name]
				make:model [name]
				make:service [name]
				make:repository [name]
				make:all [name]
				make:migration [name]`)
		}
		for _,flag:= range args{
			if flag == "migration"{
				fmt.Println("make migration")
			}
		}
	},
}

func MakeCommand() *cobra.Command {
	return make
}