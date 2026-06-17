package migrations

import (

    "github.com/beyond3800/hawk/core/migration"
	"github.com/beyond3800/hawk/core/database"
)

type LikesMigration struct{}

func (Likes LikesMigration) Up() error {
    _, err := database.HawkDB().Conn.Exec(`
        CREATE TABLE Likes (
            id BIGINT AUTO_INCREMENT PRIMARY KEY,
            name VARCHAR(255),
            email VARCHAR(255)
        )
    `)

    return err
}

func (Likes LikesMigration) Down() error {
    _, err := database.HawkDB().Conn.Exec(`
        DROP TABLE Likes
    `)

    return err
}

func init (){
    migration.Register("20260615072400_create_likes_table",LikesMigration{})
}