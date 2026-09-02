package migration

import (
	"fmt"

	"github.com/beyond3800/hawk/database"
)

func allTables() ([]string, error) {
	rows, err := database.HawkDB().Conn.Query(`
		SHOW TABLES
	`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tables []string

	for rows.Next() {
		var name string

		if err := rows.Scan(&name); err != nil {
			return nil, err
		}

		if name == "migrations" {
			continue
		}

		tables = append(tables, name)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tables, nil
}

func disableForeignKeyChecks() error {
	_, err := database.HawkDB().Conn.Exec(`
		SET FOREIGN_KEY_CHECKS = 0
	`)

	return err
}

func enableForeignKeyChecks() error {
	_, err := database.HawkDB().Conn.Exec(`
		SET FOREIGN_KEY_CHECKS = 1
	`)

	return err
}

func clearMigrations() error {
	_, err := database.HawkDB().Conn.Exec(`
		DELETE FROM migrations
	`)

	return err
}

func Fresh() error {
	if err := disableForeignKeyChecks(); err != nil {
		return err
	}

	defer enableForeignKeyChecks()

	tables, err := allTables()
	if err != nil {
		return err
	}

	for _, name := range tables {
		fmt.Println("Dropping "+name+" table 🔃")
		table := database.HawkDB().Schema().Table(name)

		if err := table.Drop().Execute(); err != nil {
			fmt.Println("Unable to drop "+name+" table ❌")
			return err
		}
	}
	fmt.Println("All table dropped")
	if err := clearMigrations(); err != nil {
		fmt.Println("Unale to clear migration ❌")
		return err
	}
	fmt.Println("Tables migration done successfully ✅")
	return Run()
}