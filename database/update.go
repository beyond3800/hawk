package database


import (
	"database/sql"
	"fmt"
	"strings"
)

func (b *Builder) Update(data map[string]any) (sql.Result, error) {

    setParts := []string{}
    values := []any{}

    for k, v := range data {
        setParts = append(setParts, k+" = ?")
        values = append(values, v)
    }

    query := fmt.Sprintf(
        "UPDATE %s SET %s",
        b.table,
        strings.Join(setParts, ", "),
    )

    if len(b.wheres) > 0 {
        query += " WHERE " + strings.Join(b.wheres, " AND ")
    }

    values = append(values, b.bindings...)

    result, sqlErr := b.exec(query, values...)
    return result, MySqlErrorFormat(sqlErr)
}