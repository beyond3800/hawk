package database

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
)

type Builder struct {
	table    string
	columns  []string
	wheres   []string
	bindings []any
	limit    int
	orMode   bool
	orderBy  string
	inserts  map[string]any
}
func (db *DB) Table(name string) *Builder {
	return &Builder{
		table: name,
	}
}

func (b *Builder) Select(columns ...string) *Builder {
	b.columns = columns
	return b
}
func (b *Builder) Where(column string, value any) *Builder {

	b.wheres = append(
		b.wheres,
		column+" = ?",
	)

	b.bindings = append(
		b.bindings,
		value,
	)

	return b
}
func (b *Builder) Limit(limit int) *Builder {
	b.limit = limit
	return b
}
func (b *Builder) First(dest any) error {

    b.limit = 1

    query, args := b.ToSQL()

    rows, err := HawkDB().Conn.Query(query, args...)
	if err != nil {
		return err
	}

	defer rows.Close()

	if !rows.Next() {
		return sql.ErrNoRows
	}

	return scanRow(rows, dest)
}
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

	// if b.orderBy != "" {
	// 	query += " ORDER BY " + b.orderBy
	// }

	if len(b.wheres) > 0 {
		query += " WHERE "
		query += strings.Join(b.wheres, " AND ")
	}

	if b.limit > 0 {
		query += fmt.Sprintf(" LIMIT %d",
			b.limit,
		)
	}
	return query, b.bindings
}
func (b *Builder) Get(dest any) error {

	query, args := b.ToSQL()

	rows, err := HawkDB().Conn.Query(query, args...)
	if err != nil {
		return err
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
func (b *Builder) OrWhere(column string, value any) *Builder {

    condition := column + " = ?"

    if len(b.wheres) == 0 {
        b.wheres = append(b.wheres, condition)
    } else {
        b.wheres = append(b.wheres, "OR "+condition)
    }

    b.bindings = append(b.bindings, value)

    return b
}
func (b *Builder) OrderBy(column, direction string) *Builder {
    b.orderBy = fmt.Sprintf("%s %s", column, direction)
    return b
}
func (b *Builder) Insert(data map[string]any) (sql.Result, error) {

    b.inserts = data

    keys := []string{}
    placeholders := []string{}
    values := []any{}

    for k, v := range data {
        keys = append(keys, k)
        placeholders = append(placeholders, "?")
        values = append(values, v)
    }

    query := fmt.Sprintf(
        "INSERT INTO %s (%s) VALUES (%s)",
        b.table,
        strings.Join(keys, ", "),
        strings.Join(placeholders, ", "),
    )

    return HawkDB().Conn.Exec(query, values...)
}
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

    return HawkDB().Conn.Exec(query, values...)
}
func (b *Builder) Delete() (sql.Result, error) {

    query := fmt.Sprintf("DELETE FROM %s", b.table)

    if len(b.wheres) > 0 {
        query += " WHERE " + strings.Join(b.wheres, " AND ")
    }

    return HawkDB().Conn.Exec(query, b.bindings...)
}
