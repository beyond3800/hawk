package database

import (
	"database/sql"

	"testing"

	"github.com/beyond3800/hawk/faker"
	_ "modernc.org/sqlite"
)

func setupTestDB(t *testing.T) *sql.DB {

	t.Helper()
    db, err := sql.Open("sqlite", ":memory:")
    if err != nil {
        t.Fatal(err)
    }

    _, err = db.Exec(`
        CREATE TABLE users(
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT,
            email TEXT
        );
    `)

    if err != nil {
        t.Fatal(err)
    }

    SetInstance(&DB{
		Conn: db,
	})
	t.Cleanup(func() {
		resetDB()
		db.Close()
	})

    return db
}

func seedUser(t *testing.T, amount int){
    t.Helper()

    for i := 0; i < amount; i++{
        _, err := HawkDB().Table("users").Create(
            map[string]any{
                "name":faker.Name(),
                "email": faker.Email(),
            },
        )
        if err != nil{
            t.Fatal(err)
        }
    }
}
func createWalletTable(t *testing.T){
    t.Helper()
    _, err := HawkDB().Conn.Exec(`
        CREATE TABLE wallets(
            id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id TEXT NOT NULL,
			balance REAL DEFAULT 0
        );
    `)

    if err != nil {
        t.Fatal(err)
    }
}

func assertCount(t *testing.T, table string, expected int) {
    t.Helper()

    count, err := HawkDB().Table(table).Count()

    if err != nil {
        t.Fatal(err)
    }

    if count != expected {
        t.Fatalf("expected %d records, got %d", expected, count)
    }
}