package database


import (
	"fmt"
)

func (b *Builder) Count() (int, error) {
	var count int

	query := fmt.Sprintf(
		"SELECT COUNT(*) FROM %s",
		b.table,
	)

	query += b.buildWhere()

	err := b.queryRow(
		query,
		b.bindings...,
	).Scan(&count)

	if err != nil {
		return 0, MySqlErrorFormat(err)
	}

	return count, nil
}