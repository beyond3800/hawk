package database

import (
	"database/sql"
	"fmt"

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

	hasSoftDeletes bool
	withTrash bool
	onlyTrash bool

	schemaErr error
}


func (b *Builder) Limit(limit int) *Builder {
	b.limit = limit
	return b
}

func (b *Builder) OrderBy(column, direction string) *Builder {
    b.orderBy = fmt.Sprintf("%s %s", column, direction)
    return b
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
