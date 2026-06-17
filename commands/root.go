package commands

import (

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "hawk",
	Short: "A Laravel-style CLI tool in Go",
	Long:  "Custom CLI built with Cobra for scaffolding, serving, and migrations.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute () error{
	RegisterCommands()
	return rootCmd.Execute()
}