package database

import (
	"database/sql"
	"fmt"

	schema "github.com/beyond3800/hawk/database/schema"
)


type DB struct {
    Conn *sql.DB
}

var instance *DB

func SetInstance(db *DB) {
	if db == nil {
        panic("attempted to register a nil database instance")
    }

    instance = db
}
func resetDB() {
    instance = nil
}

func HawkDB() *DB {
	return instance
}

func (db *DB) Schema() *schema.Schema {
	return &schema.Schema{
		Db: db.Conn,
	}
}

func (db *DB) Model(model any) (*ModelQuery, error) {
	m, ok := model.(interface {
		TableName() string
	})
	if !ok {
		return nil, fmt.Errorf("model must implement TableName() string")
	}
	return &ModelQuery{
		model:   model,
		builder: db.Table(m.TableName()),
	}, nil
}

