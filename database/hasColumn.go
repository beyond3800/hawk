package database


func (db *DB) HasColumn(table, column string) (bool, error) {
	var count int

	err := db.Conn.QueryRow(`
		SELECT COUNT(*)
		FROM information_schema.columns
		WHERE table_schema = DATABASE()
		  AND table_name = ?
		  AND column_name = ?
	`, table, column).Scan(&count)

	if err != nil {
		return false, err
	}

	return count > 0, nil
}