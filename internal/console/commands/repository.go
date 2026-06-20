package commands

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/beyond3800/hawk/lib"
	"github.com/spf13/cobra"
)

var repository = &cobra.Command{
	Use: "repository",
	Short: "Use to create repository file",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dir :="app/Http/Repository"
		path :="app/Http/Repository"
		_,err := os.Stat(dir)
		if err !=nil{
			if err := os.Mkdir(dir, 0755); err != nil{
				fmt.Println("Unable to create file")
			}
		}
		name := args[0]
		if !strings.HasSuffix(name,"Repository"){
			name = name + "Repository"
		}
		if lib.FileExist(dir,name){
				fmt.Println("This file already in" + dir)
				return
			}
		err = lib.GenerateTemplate(name,"repository",path)
		if err != nil{
			fmt.Println("Unable to create file")
			return
		}
		log.Printf(`Repository created successfully`)
	},
}



func RepositoryCommand() *cobra.Command {
	return repository
}