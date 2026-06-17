package commands

import "github.com/spf13/cobra"

var all = &cobra.Command{
	Use: "all [name]",
	Short: "Use to make all service",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		controller.Run(cmd,args)
		repository.Run(cmd,args)
		makeModelCmd.Run(cmd,args)
		service.Run(cmd,args)
	},
}




func AllCommand() *cobra.Command {
	return all
}