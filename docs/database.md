# Database

Hawk supports MySQL and PostgreSQL connections.

---

## Environment Configuration

```env
DB_ENABLED=true

DB_CONNECTION=mysql
DB_HOST=127.0.0.1
DB_PORT=3306
DB_DATABASE=hawk
DB_USERNAME=root
DB_PASSWORD=
```

---

## Connecting

Database connections are initialized automatically during application startup.

```go
bootstrap.Bootstrap()
```

---

## Accessing the Database

```go
db := database.HawkDB()
```

---

## Execute Raw Queries

```go
_, err := database.HawkDB().Conn.Exec(`
    CREATE TABLE users (
        id BIGINT AUTO_INCREMENT PRIMARY KEY,
        name VARCHAR(255)
    )
`)
```

---

## Query Builder

```go
database.Table("users").
    Where("id", 1).
    First(&user)
```

See the Query Builder documentation for additional examples.
