package migrations

import (

    "github.com/beyond3800/hawk/core/migration"
	"github.com/beyond3800/hawk/core/database"
)

type LoggingMigration struct{}

func (Logging LoggingMigration) Up() error {
    _, err := database.HawkDB().Conn.Exec(`
        CREATE TABLE Logging (
            id BIGINT AUTO_INCREMENT PRIMARY KEY,
            name VARCHAR(255),
            email VARCHAR(255)
        )
    `)

    return err
}

func (Logging LoggingMigration) Down() error {
    _, err := database.HawkDB().Conn.Exec(`
        DROP TABLE Logging
    `)

    return err
}

func init (){
    migration.Register("20260617002628_create_logging_table",LoggingMigration{})
}