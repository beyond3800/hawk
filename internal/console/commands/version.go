package commands

import (
	"fmt"

	"github.com/beyond3800/hawk"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display the version of hawk",
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string){
		fmt.Println("Hawk version:", hawk.Version())
	},
}

func VersionCommand() *cobra.Command {
	return versionCmd
}