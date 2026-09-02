package migration

import "github.com/beyond3800/hawk/database"

func allMigration() ([]string, error) {
	rows, err := database.HawkDB().Conn.Query(`
		SELECT migration 
		FROM migrations
		ORDER BY id DESC
	`)
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
	if rows.Err() != nil{
		return  nil, rows.Err()
	}

	return names, nil
}

func Reset() error{
	names, err := allMigration()

	if err != nil {
		return err
	}
	for _, name := range names {
		if err = RemoveMigration(name); err != nil {
			return err
		}
	}
	return nil
}