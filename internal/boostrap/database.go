package bootstrap

import (
	"database/sql"
	"sync"

	"github.com/beyond3800/hawk/core/database"
)

var (
	db   *sql.DB
	once sync.Once
)

func DB() *sql.DB {
	once.Do(func() {
		database.ConnectDatabase()
	})
	return db
}