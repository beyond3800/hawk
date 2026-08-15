package database

import (
	"fmt"
	"reflect"
)


func (b *Builder) Get(dest any) error {

	query, args := b.ToSQL()

	rows, err := b.query(query, args...)
	if err != nil {
		return MySqlErrorFormat(err)
	}
	defer rows.Close()

	v := reflect.ValueOf(dest)

	if v.Kind() != reflect.Ptr {
		return fmt.Errorf("destination must be pointer to slice")
	}

	sliceValue := v.Elem()

	if sliceValue.Kind() != reflect.Slice {
		return fmt.Errorf("destination must be pointer to slice")
	}

	elemType := sliceValue.Type().Elem()

	for rows.Next() {

		elem := reflect.New(elemType)

		if err := scanRow(rows, elem.Interface()); err != nil {
			return err
		}

		sliceValue.Set(
			reflect.Append(
				sliceValue,
				elem.Elem(),
			),
		)
	}

	return rows.Err()
}