package db


import (

	"github.com/beyond3800/hawk/internal/bootstrap"
	"github.com/spf13/cobra"
)



var seedCmd  = &cobra.Command{
	Use: "seed",
	Short: "Seeds the database with data",
	RunE: func(cmd *cobra.Command, args []string) error{
		return bootstrap.RunApp("hawk_seed")
	},
}


func SeedCommand() *cobra.Command {
	return seedCmd
}