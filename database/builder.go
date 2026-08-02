package database

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"

)

type Builder struct {
	db	     *DB
	table    string
	columns  []string
	wheres   []string
	bindings []any
	limit    int
	orMode   bool
	orderBy  string
	inserts  any
	offset   int

	perPage  int
	page     int

	tx	     *sql.Tx
}
func (b *Builder) query(query string, args ...any) (*sql.Rows, error) {
	if b.tx != nil {
		return b.tx.Query(query, args...)
	}

	return b.db.Conn.Query(query, args...)
}
func (b *Builder) queryRow(query string, args ...any) *sql.Row {
	if b.tx != nil {
		return b.tx.QueryRow(query, args...)
	}

	return b.db.Conn.QueryRow(query, args...)
}

func (b *Builder) exec(query string, args ...any) (sql.Result, error) {
	if b.tx != nil {
		return b.tx.Exec(query, args...)
	}

	return b.db.Conn.Exec(query, args...)
}

func (db *DB) Table(name string) *Builder {
	return &Builder{
		db: db,
		table: name,
	}
}
func (b *Builder) Table(name string) *Builder {
	return &Builder{
		db:    b.db,
		tx:    b.tx,
		table: name,
	}
}
func (db *DB) Transaction(fn func(bd *Builder) error) error {
	tx, err := db.Conn.Begin()
	if err != nil {
		return err
	}

	builder := &Builder{
		db: db,
		tx: tx,
	}

	if err := fn(builder); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("transaction error: %v, rollback error: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
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

    rows, err := b.query(query, args...)
	if err != nil {
		return MySqlErrorFormat(err)
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
	if b.offset > 0 {
		query += fmt.Sprintf(" OFFSET %d", b.offset)
	}
	return query, b.bindings
}
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
		// fieldRequired := field.Tag.Get("validate")
		// requiredSlice := strings.Split(fieldRequired, "|")
		// for _,v := range requiredSlice{
		// 	if v == "required"{
				keys = append(keys, field.Tag.Get("db"))
				placeholders = append(placeholders, "?")
				values = append(values, val.Field(i).Interface())
		// 	}
		// }
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
func (b *Builder) Delete() (sql.Result, error) {

    query := fmt.Sprintf("DELETE FROM %s", b.table)

    if len(b.wheres) > 0 {
        query += " WHERE " + strings.Join(b.wheres, " AND ")
    }

	result, sqlErr := b.exec(query, b.bindings...)
    return result, MySqlErrorFormat(sqlErr)
}
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

func (b *Builder) Offset(offset int) *Builder{
	b.offset = offset
	return b
}
func (b *Builder) Page(page int) *Builder{
	b.page = page
	return b
}
func (b *Builder) PerPage(perPage int) *Builder{
	b.perPage = perPage
	return b
}
func (b *Builder) Paginate(page, perPage int, dest any) (*Pagination, error) {

	if page < 1 {
		page = 1
	}

	if perPage < 1 {
		perPage = 15
	}

	countBuilder := b.Table(b.table)
	countBuilder.wheres = append([]string{}, b.wheres...)
	countBuilder.bindings = append([]any{}, b.bindings...)

	total, err := countBuilder.Count()
	if err != nil {
		return nil, err
	}

	b.page = page
	b.perPage = perPage
	b.limit = perPage
	b.offset = (page - 1) * perPage

	if err := b.Get(dest); err != nil {
		return nil, err
	}

	lastPage := total / perPage
	if total%perPage != 0 {
		lastPage++
	}

	return &Pagination{
		Data: dest,
		Meta: Meta{
			Page: page,
			PerPage: perPage,
			Total: total,
			LastPage: lastPage,
		},
	}, nil
}