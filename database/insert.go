package database

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
)


func (b *Builder) Insert(data any) (sql.Result, error) {

    b.inserts = data

    keys := []string{}
    placeholders := []string{}
    values := []any{}
    
	val := reflect.ValueOf(data)
	if val.Kind() == reflect.Ptr{
        val = val.Elem()
    }
	structType := val.Type()
	for i := 0; i < structType.NumField(); i++ {
        field := structType.Field(i)
			keys = append(keys, field.Tag.Get("db"))
			placeholders = append(placeholders, "?")
			values = append(values, val.Field(i).Interface())
	}
    query := fmt.Sprintf(
        "INSERT INTO %s (%s) VALUES (%s)",
        b.table,
        strings.Join(keys, ", "),
        strings.Join(placeholders, ", "),
    )
	result, sqlErr := b.exec(query, values...)
    return result, MySqlErrorFormat(sqlErr)
}