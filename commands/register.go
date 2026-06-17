package commands

import "github.com/spf13/cobra"

func RegisterRootCommand(cmd *cobra.Command) {
	rootCmd.AddCommand(cmd)
}
func RegisterMakeCommand(cmd *cobra.Command) {
	make.AddCommand(cmd)
}