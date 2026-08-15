package database

import "database/sql"

func (b *Builder) First(dest any) error {

	b.limit = 1

	query, args := b.ToSQL()

	rows, err := b.query(query, args...)
	if err != nil {
		return MySqlErrorFormat(err)
	}

	defer rows.Close()

	if !rows.Next() {
		return sql.ErrNoRows
	}

	return scanRow(rows, dest)
}