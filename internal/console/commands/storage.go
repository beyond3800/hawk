package commands

import (
	"fmt"
	"os"

	"github.com/beyond3800/hawk/internal/env"
	"github.com/spf13/cobra"
)

var storageCmd = &cobra.Command{
	Use:   "storage",
	Short: "Storage is use to create storage configuration for the application",
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {

		parser := env.New(".env")
		parser.Load()

		if parser.Has("STORAGE_ENABLED"){
			fmt.Println("Storage has been configured for this app ")
			return
		}

		if err := os.MkdirAll("storage", 0755); err != nil {
			fmt.Println(err)
			return
		}
		parser.Set("STORAGE_ENABLED", "true")
		parser.Set("STORAGE_DRIVER", "local")
		parser.Set("STORAGE_ROOT", "storage")
		parser.Set("STORAGE_URL", "/storage")
		parser.Save()

		fmt.Println("✔ Storage configured successfully.")
	},
}

func StorageCommand() *cobra.Command {
	return storageCmd
}