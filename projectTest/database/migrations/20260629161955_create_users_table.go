package Migrations

import (

    "github.com/beyond3800/hawk/migration"
	"github.com/beyond3800/hawk/database"
)

type UsersMigration struct{}

func (Users UsersMigration) Up() error {
    _, err := database.HawkDB().Conn.Exec(`
        CREATE TABLE users (
            id BIGINT AUTO_INCREMENT PRIMARY KEY,
            name VARCHAR(255),
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
        )
    `)

    return err
}

func (Users UsersMigration) Down() error {
    _, err := database.HawkDB().Conn.Exec(`
        DROP TABLE users
    `)

    return err
}

func init (){
    migration.Register("20260629161955_create_users_table",UsersMigration{})
}