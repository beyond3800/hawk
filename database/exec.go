package database


import (
	"database/sql"
)

func (b *Builder) exec(query string, args ...any) (sql.Result, error) {
	if b.tx != nil {
		return b.tx.Exec(query, args...)
	}

	return b.db.Conn.Exec(query, args...)
}