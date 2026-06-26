# Migrations

Migrations allow you to create and modify database tables using Go code.

---

## Create a Migration

Generate a new migration:

```bash
hawk make migration create_users_table
```

This creates a migration file inside:

```text
database/migrations/
```

---

## Creating a Migration

```go
type UsersMigration struct{}

func (UsersMigration) Up() error {
    _, err := database.HawkDB().Conn.Exec(`
        CREATE TABLE users (
            id BIGINT AUTO_INCREMENT PRIMARY KEY,
            name VARCHAR(255),
            email VARCHAR(255)
        )
    `)

    return err
}

func (UsersMigration) Down() error {
    _, err := database.HawkDB().Conn.Exec(`
        DROP TABLE users
    `)

    return err
}
```

---

## Registering the Migration

Register the migration inside the `init` function.

```go
func init() {
    migration.Register(
        "20260617101109_create_users_table",
        UsersMigration{},
    )
}
```

---

## Run Migrations

Execute all pending migrations:

```bash
hawk migrate
```

---

## Roll Back Migrations

Revert the latest migration batch:

```bash
hawk rollback
```

---

## Migration Status

Display migration status:

```bash
hawk status
```

---

## Example Output

```text
✓ 20260617101109_create_users_table
Database migrated successfully
```
