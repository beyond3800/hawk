package database



func (db *DB) hasSoftDeletes(table string) (bool, error) {
	hasColumn, err := db.HasColumn(table, "deleted_at")
	if err != nil {
		return false, err
	}

	return hasColumn, nil
}