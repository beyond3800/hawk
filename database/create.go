package database

import (
	"database/sql"
	"fmt"
	"strings"
)

func (b *Builder) Create(data map[string]any) (sql.Result, error) {

	if len(data) == 0 {
		return nil, fmt.Errorf("no data supplied")
	}

	columns := make([]string, 0, len(data))
	placeholders := make([]string, 0, len(data))
	values := make([]any, 0, len(data))

	for column, value := range data {
		columns = append(columns, column)
		placeholders = append(placeholders, "?")
		values = append(values, value)
	}

	query := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s)",
		b.table,
		strings.Join(columns, ", "),
		strings.Join(placeholders, ", "),
	)

	result, sqlErr := b.exec(query, values...)
	return result, MySqlErrorFormat(sqlErr)
}