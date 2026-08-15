package Make

import "github.com/spf13/cobra"

var all = &cobra.Command{
	Use: "all [name]",
	Short: "Use to make all service",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		controller.Run(cmd,args)
		repository.Run(cmd,args)
		modelCmd.Run(cmd,args)
		serviceCmd.Run(cmd,args)
	},
}




func AllCommand() *cobra.Command {
	return all
}