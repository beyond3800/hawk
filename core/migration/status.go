package migration

import (
	"fmt"

	"github.com/beyond3800/hawk/core/database"
)

func ExecutedMigrations() (map[string]bool, error) {

	rows, err := database.HawkDB().Conn.Query(
		"SELECT migration FROM migrations",
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	executed := make(map[string]bool)

	for rows.Next() {
		var name string
		rows.Scan(&name)

		executed[name] = true
	}

	return executed, nil
}


func Status() error {

	executed, err := ExecutedMigrations()
	if err != nil {
		return err
	}

	for _, migration := range migrations {

		status := "[ ]"

		if executed[migration.Name] {
			status = "[X]"
		}

		fmt.Printf("%s %s\n",
			status,
			migration.Name,
		)
	}

	return nil
}