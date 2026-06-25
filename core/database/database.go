package database

import (
	"database/sql"
)


type DB struct {
    Conn *sql.DB
}

var instance *DB

func SetInstance(db *DB) {
	instance = db
}

func HawkDB() *DB {
	return instance
}