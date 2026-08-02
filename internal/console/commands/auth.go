package commands

import (
	"fmt"

	"github.com/beyond3800/hawk/internal/env"
	"github.com/beyond3800/hawk/util"
	"github.com/spf13/cobra"
)

var auth = &cobra.Command{
	Use:   "auth",
	Short: "Auth is use to create authentication secret key and issuer",
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		parser := env.New(".env")
		parser.Load()
		if parser.Has("APP_SECRET_KEY"){
			fmt.Println("Auth secret has been created for this app ")
			return
		}
		generatedSecret, err := util.GenerateSecret()
		if err != nil{
			fmt.Println("Unable to generate auth secret")
		}

		parser.Set("APP_SECRET_KEY",generatedSecret)
		parser.Set("APP_ISSUER","Hawk")
		parser.Save()
		fmt.Println("Auth created successfully for this App")
	},
}

func AuthCommand() *cobra.Command {
	return auth
}