package database

import (
	"database/sql"
	"fmt"
	"reflect"
)

func scanRow(rows *sql.Rows, dest any) error {

	v := reflect.ValueOf(dest)

	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("destination must be a pointer to struct")
	}

	structValue := v.Elem()
	structType := structValue.Type()

	columns, err := rows.Columns()
	if err != nil {
		return err
	}

	fieldMap := make(map[string]reflect.Value)

	for i := 0; i < structType.NumField(); i++ {

		field := structType.Field(i)

		tag := field.Tag.Get("db")

		if tag == "" {
			tag = field.Name
		}

		fieldMap[tag] = structValue.Field(i)
	}

	scanTargets := make([]any, len(columns))

	for i, column := range columns {

		if field, ok := fieldMap[column]; ok {

			scanTargets[i] = field.Addr().Interface()

		} else {

			var dummy any
			scanTargets[i] = &dummy
		}
	}

	return rows.Scan(scanTargets...)
}