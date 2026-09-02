package database

import (
	"database/sql"
	"fmt"
	"strings"
)



func (b *Builder) Delete() (sql.Result, error) {
	if b.schemaErr != nil && b.hasSoftDeletes {
		return nil, b.schemaErr
	}

	if b.hasSoftDeletes {
		query := fmt.Sprintf(
			"UPDATE %s SET deleted_at = CURRENT_TIMESTAMP",
			b.table,
		)

		if len(b.wheres) > 0 {
			query += " WHERE " + strings.Join(b.wheres, " AND ")
		}

		result, err := b.exec(query, b.bindings...)

		return result, MySqlErrorFormat(err)
	}

	return b.ForceDelete()
}
func (b *Builder) ForceDelete() (sql.Result, error) {
    

    query := fmt.Sprintf("DELETE FROM %s", b.table)

    if len(b.wheres) > 0 {
        query += " WHERE " + strings.Join(b.wheres, " AND ")
    }

	result, sqlErr := b.exec(query, b.bindings...)
    return result, MySqlErrorFormat(sqlErr)
}