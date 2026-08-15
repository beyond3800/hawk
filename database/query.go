package database


import (
	"database/sql"
)


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