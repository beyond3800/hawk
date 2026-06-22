package database

import (
	"database/sql"
    _ "github.com/go-sql-driver/mysql"
)

type DB struct {
    Conn *sql.DB
}

var instance *DB

func HawkDB() *DB {
    return instance
}