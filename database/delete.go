package database

import (
	"database/sql"
	"fmt"
	"strings"
)


func (b *Builder) Delete() (sql.Result, error) {

    query := fmt.Sprintf("DELETE FROM %s", b.table)

    if len(b.wheres) > 0 {
        query += " WHERE " + strings.Join(b.wheres, " AND ")
    }

	result, sqlErr := b.exec(query, b.bindings...)
    return result, MySqlErrorFormat(sqlErr)
}