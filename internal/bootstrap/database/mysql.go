package database

import (
	"database/sql"
	"fmt"
	"sync"

	"github.com/beyond3800/hawk/database"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db       *sql.DB
	once     sync.Once
)

func connectMySQL(config Config) error {
	var connectErr error
	once.Do(func() {

		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?parseTime=true",
			config.User,
			config.Password,
			config.Host,
			config.Port,
			config.Database,
		)

		db, connectErr = sql.Open("mysql", dsn)
		if connectErr != nil {
			connectErr = fmt.Errorf("failed to connect to database: %w", connectErr)
			return
		}

		if connectErr = db.Ping(); connectErr != nil {
			connectErr = fmt.Errorf("database unreachable: %w", connectErr)
			return
		}
		database.SetInstance(&database.DB{
			Conn: db,
		})

	})

	return connectErr
}
