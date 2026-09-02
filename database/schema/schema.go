package database

import (
	"database/sql"
)

type Schema struct {
	Db *sql.DB
}