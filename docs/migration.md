
---

# `docs/migrations.md`

```md
# Migrations

## Create Migration

```bash
hawk make migration create_users_table


type UsersMigration struct{}

func (Users UsersMigration) Up() error {
    _, err := database.HawkDB().Conn.Exec(`
        CREATE TABLE Users (
            id BIGINT AUTO_INCREMENT PRIMARY KEY,
            name VARCHAR(255),
            email VARCHAR(255)
        )
    `)

    return err
}

func (Users UsersMigration) Down() error {
    _, err := database.HawkDB().Conn.Exec(`
        DROP TABLE Users
    `)

    return err
}

func init (){
    migration.Register("20260617101109_create_users_table",UsersMigration{})
}