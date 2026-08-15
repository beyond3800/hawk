package database

import (
	"database/sql"
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