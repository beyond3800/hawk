package database

import (
	"database/sql"
	"fmt"
)


type DB struct {
    Conn *sql.DB
}

var instance *DB

func SetInstance(db *DB) {
	if db == nil {
        fmt.Println("attempted to register a nil database instance")
    }

    instance = db
	
}

func HawkDB() *DB {
	return instance
}