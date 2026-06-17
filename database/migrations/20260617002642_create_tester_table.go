package migrations

import (

    "github.com/beyond3800/hawk/core/migration"
	"github.com/beyond3800/hawk/core/database"
)

type TesterMigration struct{}

func (Tester TesterMigration) Up() error {
    _, err := database.HawkDB().Conn.Exec(`
        CREATE TABLE Tester (
            id BIGINT AUTO_INCREMENT PRIMARY KEY,
            name VARCHAR(255),
            email VARCHAR(255)
        )
    `)

    return err
}

func (Tester TesterMigration) Down() error {
    _, err := database.HawkDB().Conn.Exec(`
        DROP TABLE Tester
    `)

    return err
}

func init (){
    migration.Register("20260617002642_create_tester_table",TesterMigration{})
}