package hawk

import (
	"fmt"
	"os"

	engine "github.com/beyond3800/hawk/engine"
	"github.com/beyond3800/hawk/internal/console/migration"
	_"github.com/beyond3800/hawk/seeder"
)

func Execute(app *engine.Hawk){
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
		case "hawk_reset":
			if err :=migration.Reset(); err != nil {
				fmt.Println("Reset error:", err)
				return
			}
			return
		case "hawk_refresh":
			if err :=migration.Refresh(); err != nil {
				fmt.Println("Reset error:", err)
				return
			}
			return
		case "hawk_fresh":
			if err :=migration.Fresh(); err != nil {
				fmt.Println("Reset error:", err)
				return
			}
			return
		default:
			return
		}

	}
	port := os.Getenv("APP_PORT")
    if port == "" {
        port = ":8080"
    }

    app.Run(port)
}