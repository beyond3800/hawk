package database


import (
	"fmt"
	"strings"
)

func (b *Builder) ToSQL() (string, []any) {
	columns := "*"

	if len(b.columns) > 0 {
		columns = strings.Join(b.columns, ", ")
	}

	query := fmt.Sprintf(
		"SELECT %s FROM %s",
		columns,
		b.table,
	)

	query += b.buildWhere()

	if b.limit > 0 {
		query += fmt.Sprintf(" LIMIT %d", b.limit)
	}

	if b.offset > 0 {
		query += fmt.Sprintf(" OFFSET %d", b.offset)
	}

	return query, b.bindings
}