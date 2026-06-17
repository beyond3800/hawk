package migrations

import (

    "github.com/beyond3800/hawk/core/migration"
	"github.com/beyond3800/hawk/core/database"
)

type PostsMigration struct{}

func (Posts PostsMigration) Up() error {
    _, err := database.HawkDB().Conn.Exec(`
        CREATE TABLE Posts (
            id BIGINT AUTO_INCREMENT PRIMARY KEY,
            name VARCHAR(255),
            email VARCHAR(255)
        )
    `)

    return err
}

func (Posts PostsMigration) Down() error {
    _, err := database.HawkDB().Conn.Exec(`
        DROP TABLE Posts
    `)

    return err
}

func init (){
    migration.Register("20260614230151_create_posts_table",PostsMigration{})
}