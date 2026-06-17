package commands

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/beyond3800/hawk/lib"
	"github.com/spf13/cobra"
)

var service = &cobra.Command{
	Use: "service",
	Short: "Use to create service file",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dir := "services"
		name := args[0]
		if !strings.HasSuffix(name,"Service"){
			name = name + "Service"
		}
		if _,err := os.Stat(dir); err != nil{
			if err := os.MkdirAll(dir,0755); err != nil{
				fmt.Println("Unable to create folder")
			}
		}
		if lib.FileExist(dir,name){
			fmt.Println("This file already in" + dir)
			return
		}
		
		err := lib.GenerateTemplate(name,"service")
		if err != nil{
			log.Fatal(err)
			log.Fatal("Unable to create service")
			return
		}
		log.Println("Service created successfully")
	},
}

func ServiceCommand() *cobra.Command {
	return service
}