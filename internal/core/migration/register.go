package migration

import (
	"fmt"
	_ "fmt"
	_ "os"

	"github.com/beyond3800/hawk/internal/core/database"
)

type RegisteredMigration struct{
	Name string
	M Migration
}

var migrations []RegisteredMigration


func Register(name string, m Migration) {
	
	migrations = append(migrations, RegisteredMigration{
		Name: name,
		M: m,
	})
}

func Run() error {
	migrationHappen := false
	if err := createMigrationTable(); err != nil {
		return err
	}

	batch,err := nextBatch()
	if err != nil {
		return err
	}
	for _, migration := range migrations {
		if hasMigrated(migration.Name){
			continue
		}
		if err := migration.M.Up(); err != nil {
			migrationHappen = true
			return err
		}
		fmt.Println(migration.Name)
		saveMigration(migration.Name,batch)
	}
	if migrationHappen{
		fmt.Println("Database migrated successfully")
	}else{
		fmt.Println("Noting to migrate")
	}

	return nil
}

func hasMigrated(name string) bool {

    var count int

    err := database.HawkDB().Conn.QueryRow(
        "SELECT COUNT(*) FROM migrations WHERE migration = ?",
        name,
    ).Scan(&count)

    if err != nil {
        return false
    }

    return count > 0
}

func nextBatch ()(int, error){
	var batch int
	err := database.HawkDB().Conn.QueryRow(
		"SELECT COALESCE(MAX(batch), 0) FROM migrations",
	).Scan(&batch)
	if err != nil {
		return 0,nil
	}
	return batch+1, nil
}

func saveMigration (migration string, batch int) error{
	_,err :=database.HawkDB().Table("migrations").Insert(map[string]any{
		"migration":migration,
		"batch":batch,
	})
	if err != nil{
		return err
	}
	return nil
}

func createMigrationTable() error{
	    _, err := database.HawkDB().Conn.Exec(`
        CREATE TABLE IF NOT EXISTS Migrations (
            id BIGINT AUTO_INCREMENT PRIMARY KEY,
            migration VARCHAR(255),
            batch INT,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        )
    `)
	return err
}

// for another update

// migrations, err := os.ReadDir("database/migrations")
// if err != nil{
// 	return err
// }
// for _, m := range migrations{
// 	if m.Name() != "" {
// 		return fmt.Errorf("Migration exist already")
// 	}

// }
