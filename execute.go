package hawk

import (
	"os"

	hawk "github.com/beyond3800/hawk/engine"
	"github.com/beyond3800/hawk/internal/console/migration"
)

func Execute(app *hawk.Hawk){
	if len(os.Args)> 1{
		switch os.Args[1] {
		case "hawk_migrate":
			migration.Run()
			return
		case "hawk_rollback":
			migration.Rollback()
			return
		case "hawk_status":
			migration.Status()
			return
		}

	}
	port := os.Getenv("APP_PORT")
    if port == "" {
        port = ":8080"
    }

    app.Run(port)
}