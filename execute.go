package hawk

import (
	"fmt"
	"os"

	hawk "github.com/beyond3800/hawk/engine"
	"github.com/beyond3800/hawk/internal/console/migration"
)

func Execute(app *hawk.Hawk){
	if len(os.Args)> 1{
		switch os.Args[1] {
		case "hawk_migrate":
			if err :=migration.Run(); err != nil {
				fmt.Println("Migration error:", err)
				return
			}
			fmt.Println("Migration completed successfully")
			return
		case "hawk_rollback":
			if err :=migration.Rollback(); err != nil {
				fmt.Println("Rollback error:", err)
				return
			}
			return
		case "hawk_status":
			if err :=migration.Status(); err != nil {
				fmt.Println("Status error:", err)
				return
			}
			return
		}

	}
	port := os.Getenv("APP_PORT")
    if port == "" {
        port = ":8080"
    }

    app.Run(port)
}