package main

import (
	"fmt"
	"os"

	"github.com/beyond3800/hawk/commands"

)

func main() {
	
	if err := commands.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}