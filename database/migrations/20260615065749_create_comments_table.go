package migrations

import (

    "github.com/beyond3800/hawk/core/migration"
	"github.com/beyond3800/hawk/core/database"
)

type CommentsMigration struct{}


func (Comments CommentsMigration) Up() error {
    _, err := database.HawkDB().Conn.Exec(`
        CREATE TABLE Comments (
            id BIGINT AUTO_INCREMENT PRIMARY KEY,
            name VARCHAR(255),
            email VARCHAR(255)
        )
    `)

    return err
}

func (Comments CommentsMigration) Down() error {
    _, err := database.HawkDB().Conn.Exec(`
        DROP TABLE Comments
    `)

    return err
}

func init (){
    migration.Register("20260615065749_create_comments_table",CommentsMigration{})
}