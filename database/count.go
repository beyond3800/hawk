package database


import (
	"fmt"
	"strings"
)

func (b *Builder) Count() (int, error) {

	var count int

	query := fmt.Sprintf(
		"SELECT COUNT(*) FROM %s",
		b.table,
	)

	if len(b.wheres) > 0 {
		query += " WHERE " + strings.Join(b.wheres, " AND ")
	}

	err := b.queryRow(query, b.bindings...).Scan(&count)
	if err != nil {
		return 0, MySqlErrorFormat(err)
	}

	return count, nil
}