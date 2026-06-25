package migration

import (
	"fmt"

	"github.com/beyond3800/hawk/core/database"
)

func LatestBatch() (int, error) {
	var batch int

	err := database.HawkDB().Conn.QueryRow(`
		SELECT COALESCE(MAX(batch), 0)
		FROM migrations
	`).Scan(&batch)

	return batch, err
}


func BatchMigrations(batch int) ([]string, error) {

	rows, err := database.HawkDB().Conn.Query(`
		SELECT migration
		FROM migrations
		WHERE batch = ?
		ORDER BY id DESC
	`, batch)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var names []string

	for rows.Next() {
		var name string

		if err := rows.Scan(&name); err != nil {
			return nil, err
		}

		names = append(names, name)
	}

	return names, nil
}

func FindMigration(name string) Migration {

	for _, migration := range migrations {

		if migration.Name == name {
			return migration.M
		}
	}

	return nil
}
func RemoveMigration(name string) error {

	_, err := database.HawkDB().Conn.Exec(`
		DELETE FROM migrations
		WHERE migration = ?
	`, name)

	return err
}
func Rollback() error{
	batch, err := LatestBatch()
	fmt.Println("Latest batch:", batch)
	if err != nil {
		return err
	}

	if batch == 0 {
		fmt.Println("Nothing to rollback")
		return nil
	}

	names, err := BatchMigrations(batch)
	if err != nil {
		return err
	}

	for _, name := range names {

		migration := FindMigration(name)

		if migration == nil {
			return fmt.Errorf(
				"migration %s not found",
				name,
			)
		}

		fmt.Println("Rolling back:", name)

		if err := migration.Down(); err != nil {
			return err
		}

		if err := RemoveMigration(name); err != nil {
			return err
		}
	}

	return nil
}
