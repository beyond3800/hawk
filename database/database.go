package database

import (
	"fmt"
	"database/sql"
)


type DB struct {
    Conn *sql.DB
}

var instance *DB

func SetInstance(db *DB) {
	instance = db
	fmt.Println(instance)
}

func HawkDB() *DB {
	return instance
}