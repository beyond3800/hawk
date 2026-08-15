# Database

Hawk provides a database layer and query builder for interacting with relational databases.

The database layer provides functionality for:

* Database connections
* Query building
* Creating records
* Retrieving records
* Counting records
* Pagination
* Transactions
* Database migrations
* Seed data

The main database API is accessed through `HawkDB()`.

---

# Database Instance

Hawk exposes the configured database through:

```go
db := HawkDB()
```

The returned database instance provides access to the query builder and database connection.

For example:

```go
users := HawkDB().
    Table("users")
```

---

# Query Builder

Hawk's query builder allows database operations to be constructed using method calls.

A typical query begins by selecting a table:

```go
query := HawkDB().Table("users")
```

You can then perform operations on the query builder.

For example:

```go
var users []User

err := HawkDB().
    Table("users").
    Get(&users)
```

---

# Creating Records

Use `Create()` to insert a new record.

```go
result, err := HawkDB().
    Table("users").
    Create(map[string]any{
        "name":  "Adam",
        "email": "adam@test.com",
    })

if err != nil {
    return err
}
```

The data is supplied as:

```go
map[string]any
```

The map keys represent database columns.

The map values represent the values inserted into those columns.

For example:

```go
map[string]any{
    "name":     "Adam",
    "email":    "adam@test.com",
    "password": "secret",
}
```

---

# Empty Create Data

`Create()` requires data.

Passing an empty map results in an error:

```go
_, err := HawkDB().
    Table("users").
    Create(map[string]any{})
```

The database operation is not executed because there is no data to insert.

---

# Insert Result

`Create()` returns an `sql.Result`.

```go
result, err := HawkDB().
    Table("users").
    Create(data)
```

This allows you to access information about the inserted record using Go's standard SQL API.

For example:

```go
id, err := result.LastInsertId()

if err != nil {
    return err
}
```

This is useful when the database generates an ID automatically.

---

# Using the Inserted ID

A common use case is creating related records.

For example:

```go
result, err := tx.Table("users").Create(map[string]any{
    "name":  "Adam",
    "email": "adam@test.com",
})

if err != nil {
    return err
}

userID, err := result.LastInsertId()

if err != nil {
    return err
}

_, err = tx.Table("wallets").Create(map[string]any{
    "user_id": userID,
    "balance": 0,
})

if err != nil {
    return err
}
```

This creates a user first and then uses the generated ID to create the user's wallet.

---

# Retrieving Records

Hawk provides `Get()` for retrieving records.

For example:

```go
var users []User

err := HawkDB().
    Table("users").
    Get(&users)

if err != nil {
    return err
}
```

`Get()` populates the destination passed to it.

A slice is useful when multiple records are expected:

```go
var users []User

err := HawkDB().
    Table("users").
    Get(&users)
```

---

# First Record

Use `First()` when you want the first matching record.

```go
var user User

err := HawkDB().
    Table("users").
    First(&user)

if err != nil {
    return err
}
```

This is useful when the application expects a single record.

---

# Count

Use `Count()` to determine how many records exist.

```go
count, err := HawkDB().
    Table("users").
    Count()

if err != nil {
    return err
}
```

For example:

```go
count, err := HawkDB().
    Table("users").
    Count()

if err != nil {
    return err
}

fmt.Println("Users:", count)
```

---

# Pagination

Hawk provides pagination through `Paginate()`.

```go
var users []User

paginated, err := HawkDB().
    Table("users").
    Paginate(1, 10, &users)

if err != nil {
    return err
}
```

The arguments represent:

```text
Paginate(page, perPage, destination)
```

For example:

```go
Paginate(1, 10, &users)
```

means:

```text
Page:    1
PerPage: 10
```

---

# Pagination Metadata

The pagination result contains metadata.

For example:

```go
paginated.Meta.Total
paginated.Meta.PerPage
```

A response can therefore be constructed using:

```go
c.JSON(200, paginated)
```

A typical pagination response can contain:

```text
Data
Meta
```

The exact structure depends on the pagination type provided by the Hawk database package.

---

# Pagination Example

```go
func Users(c *hawk.Context) {

    page := 1
    perPage := 10

    var users []User

    result, err := HawkDB().
        Table("users").
        Paginate(page, perPage, &users)

    if err != nil {
        c.JSON(500, map[string]any{
            "error": "Unable to retrieve users",
        })
        return
    }

    c.JSON(200, result)
}
```

---

# Transactions

Transactions allow multiple database operations to be treated as one unit.

Use:

```go
HawkDB().Transaction(...)
```

For example:

```go
err := HawkDB().Transaction(func(tx *Builder) error {

    // database operations

    return nil
})
```

If the callback returns `nil`, the transaction can be committed.

If the callback returns an error, the transaction is rolled back.

---

# Transaction Example

Suppose an application creates a user and a wallet.

Both operations should succeed together.

```go
err := HawkDB().Transaction(func(tx *Builder) error {

    result, err := tx.Table("users").Create(map[string]any{
        "name":  "Adam",
        "email": "adam@test.com",
    })

    if err != nil {
        return err
    }

    userID, err := result.LastInsertId()

    if err != nil {
        return err
    }

    _, err = tx.Table("wallets").Create(map[string]any{
        "user_id": userID,
        "balance": 0,
    })

    if err != nil {
        return err
    }

    return nil
})
```

---

# Transaction Commit

If every operation succeeds:

```go
return nil
```

The transaction can be committed.

Conceptually:

```text
BEGIN
  ↓
Create User
  ↓
Create Wallet
  ↓
Success
  ↓
COMMIT
```

---

# Transaction Rollback

If an operation fails:

```go
return err
```

the transaction should be rolled back.

Conceptually:

```text
BEGIN
  ↓
Create User
  ↓
Create Wallet
  ↓
Error
  ↓
ROLLBACK
```

This means the earlier successful operations are also undone.

For example, if the user is created successfully but wallet creation fails, the user should not remain in the database.

---

# Why Transactions Matter

Without a transaction:

```text
Create User     ✓
Create Wallet   ✗
```

the database could contain an incomplete state.

With a transaction:

```text
Create User     ✓
Create Wallet   ✗
       ↓
ROLLBACK
```

both operations are undone.

Transactions are especially useful for:

* Wallet operations
* Payments
* Orders
* Account creation
* Transfers
* Inventory updates
* Creating related records

---

# Builder and Transactions

Inside a transaction, use the transaction builder passed to the callback:

```go
func(tx *Builder) error
```

For example:

```go
err := HawkDB().Transaction(func(tx *Builder) error {

    _, err := tx.Table("users").Create(data)

    if err != nil {
        return err
    }

    return nil
})
```

Use `tx` for operations that should be part of that transaction.

Avoid switching back to `HawkDB()` for an operation that is supposed to participate in the same transaction.

---

# Database Errors

Database operations can return errors.

Always check the error:

```go
result, err := HawkDB().
    Table("users").
    Create(data)

if err != nil {
    return err
}
```

Similarly:

```go
var users []User

err := HawkDB().
    Table("users").
    Get(&users)

if err != nil {
    return err
}
```

Ignoring database errors can result in incorrect application behavior.

---

# Error Formatting

Hawk can format database-specific errors into application-level errors.

For example, database errors such as:

```text
Duplicate key
Foreign key constraint failed
Column cannot be null
Data too long
```

can be converted into simpler errors for application code.

This allows application code to avoid depending directly on every database-specific error code.

---

# Database Connections

The Hawk database layer is backed by Go's SQL database functionality.

The database connection is represented by:

```go
*sql.DB
```

The database wrapper contains the connection:

```go
type DB struct {
    Conn *sql.DB
}
```

The connection can be accessed when lower-level database functionality is required.

For example:

```go
db := HawkDB()

rows, err := db.Conn.Query(
    "SELECT * FROM users",
)
```

Most application code should prefer Hawk's query builder when the required operation is supported.

---

# Direct SQL

Hawk's database connection can be used directly when the query builder does not provide the functionality required by an application.

For example:

```go
rows, err := HawkDB().Conn.Query(
    "SELECT id, name FROM users WHERE email = ?",
    email,
)
```

Direct SQL should be used carefully.

Always use parameterized queries for values supplied by users:

```go
rows, err := HawkDB().Conn.Query(
    "SELECT * FROM users WHERE email = ?",
    email,
)
```

Avoid constructing SQL by concatenating user input.

Do not do:

```go
query := "SELECT * FROM users WHERE email = '" + email + "'"
```

Parameterized queries help prevent SQL injection.

---

# Database Testing

Hawk's database layer can be tested using a test database.

A typical test setup creates a database connection before running database tests:

```go
func TestCreate(t *testing.T) {

    setupTestDB(t)

    _, err := HawkDB().
        Table("users").
        Create(map[string]any{
            "name":  "Adam",
            "email": "adam@test.com",
        })

    if err != nil {
        t.Fatal(err)
    }
}
```

The test can then verify the database state:

```go
count, err := HawkDB().
    Table("users").
    Count()

if err != nil {
    t.Fatal(err)
}

if count != 1 {
    t.Fatalf(
        "expected 1 user, got %d",
        count,
    )
}
```

---

# Testing Transactions

Transactions should be tested for both successful and failed operations.

A successful transaction:

```go
func TestTransaction(t *testing.T) {

    setupTestDB(t)
    createWalletTable(t)

    err := HawkDB().Transaction(func(tx *Builder) error {

        result, err := tx.Table("users").Create(map[string]any{
            "name":  "Adam",
            "email": "adam@test.com",
        })

        if err != nil {
            return err
        }

        userID, err := result.LastInsertId()

        if err != nil {
            return err
        }

        _, err = tx.Table("wallets").Create(map[string]any{
            "user_id": userID,
            "balance": 0,
        })

        return err
    })

    if err != nil {
        t.Fatal(err)
    }
}
```

A rollback test should intentionally produce an error and verify that previous operations were undone.

```go
if err == nil {
    t.Fatal("expected transaction to fail")
}
```

Then verify the database:

```go
assertCount(t, "users", 0)
assertCount(t, "wallets", 0)
```

---

# Pagination Testing

Pagination should also be tested using enough data to verify that the query actually paginates.

For example, create 100 users:

```go
seedUser(t, 100)
```

Then request the first page:

```go
var users []User

paginated, err := HawkDB().
    Table("users").
    Paginate(1, 10, &users)

if err != nil {
    t.Fatal(err)
}
```

Verify the metadata:

```go
if paginated.Meta.Total != 100 {
    t.Fatalf(
        "expected total 100, got %d",
        paginated.Meta.Total,
    )
}

if paginated.Meta.PerPage != 10 {
    t.Fatalf(
        "expected per page 10, got %d",
        paginated.Meta.PerPage,
    )
}
```

Pagination tests should also verify the number of records returned for the requested page.

---

# Recommended Database Workflow

A typical Hawk database workflow is:

```text
Create Migration
      ↓
Run Migration
      ↓
Create Seeder / Test Data
      ↓
Query Database
      ↓
Validate Results
      ↓
Write Tests
```

For application development:

```text
Migration
   ↓
Model / Struct
   ↓
Query Builder
   ↓
Application Logic
   ↓
Transaction when required
   ↓
Response
```

---

# Database API Summary

The main database operations currently documented are:

| API             | Purpose                             |
| --------------- | ----------------------------------- |
| `HawkDB()`      | Get the configured Hawk database    |
| `Table()`       | Select a database table             |
| `Create()`      | Insert a record                     |
| `Get()`         | Retrieve records                    |
| `First()`       | Retrieve the first record           |
| `Count()`       | Count records                       |
| `Paginate()`    | Retrieve paginated records          |
| `Transaction()` | Execute operations in a transaction |

---

# Example: Complete Database Operation

A complete route might look like:

```go
router.Post("/users", func(c *hawk.Context) {

    var input User

    err := json.NewDecoder(
        c.Request.Body,
    ).Decode(&input)

    if err != nil {
        c.JSON(400, map[string]any{
            "error": "Invalid request",
        })
        return
    }

    result, err := HawkDB().
        Table("users").
        Create(map[string]any{
            "name":  input.Name,
            "email": input.Email,
        })

    if err != nil {
        c.JSON(500, map[string]any{
            "error": "Unable to create user",
        })
        return
    }

    id, err := result.LastInsertId()

    if err != nil {
        c.JSON(500, map[string]any{
            "error": "Unable to retrieve user ID",
        })
        return
    }

    c.JSON(201, map[string]any{
        "id": id,
    })
})
```

The flow is:

```text
HTTP Request
     ↓
Decode JSON
     ↓
Validate Input
     ↓
Create Database Record
     ↓
Get Inserted ID
     ↓
Return JSON
```

---

# Summary

Hawk's database layer provides a query builder on top of Go's SQL functionality.

The main workflow is:

```go
HawkDB().
    Table("users").
    Create(data)
```

and:

```go
HawkDB().
    Table("users").
    Get(&users)
```

For single records:

```go
HawkDB().
    Table("users").
    First(&user)
```

For counts:

```go
HawkDB().
    Table("users").
    Count()
```

For pagination:

```go
HawkDB().
    Table("users").
    Paginate(1, 10, &users)
```

For multiple related operations:

```go
HawkDB().Transaction(func(tx *Builder) error {
    // database operations
    return nil
})
```

Hawk's database layer is designed to make common database operations concise while still allowing applications to access the underlying Go SQL connection when necessary.
