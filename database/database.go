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
	fmt.Println(db)
	if instance != nil {
		instance = db
	}
	
}

func HawkDB() *DB {
	return instance
}