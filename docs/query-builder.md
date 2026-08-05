# Query Builder

Hawk Query Builder

The Hawk Query Builder provides a fluent, expressive interface for interacting with your database. It supports retrieving, inserting, updating, deleting, pagination, and aggregation while keeping your code clean and readable.

---

## Retrieving All Records

```go
var users []User

err := database.Table("users").
    Get(&users)
```

---

## Retrieving a Single Record

```go
var user User

err := database.Table("users").
    Where("id", 1).
    First(&user)
```

---

## Where Clauses

```go
database.Table("users").
    Where("name", "John")
```

---

## Or Where Clauses

```go
database.Table("users").
    Where("name", "John").
    OrWhere("name", "Jane")
```

---

## Ordering Results

```go
database.Table("users").
    OrderBy("created_at", "DESC").
    Get(&users)
```

You may use:

* `ASC`
* `DESC`

---

## Inserting Records

```go
Insert

The Insert method inserts a record using a struct. Hawk automatically maps the exported struct fields to their corresponding database columns using db.

user := User{
    ID          string        `json:"id" db:"id"`
    Name        string        `json:"name" validate:"required" db:"name"`
    Username    string        `json:"username" validate:"required" db:"username"`
    Email       string        `json:"email" validate:"required|email" db:"email"`
    Password    string         `json:"password" validate:"required|min:6" db:"password"`
    Created_at  time.Time     `json:"created_at"`
    Updated_at  sql.NullTime  `json:"updated_at"`
}

err := database.
    Table("users").
    Insert(&user)

Use Insert when working with models or strongly typed data.
```

---

## Create Records

The Create method inserts a record using a map[string]any.

err := database.
    Table("users").
    Create(map[string]any{
        "name":  "Adam",
        "email": "adam@example.com",
    })

Use Create when building records dynamically or when the fields are not known at compile time.
```

---

## Updating Records

```go
database.Table("users").
    Where("id", 1).
    Update(map[string]any{
        "name": "Updated Name",
    })
```

---

## Deleting Records

```go
database.Table("users").
    Where("id", 1).
    Delete()
```

---

## Chaining Queries

Queries can be chained together.

```go
database.Table("users").
    Where("status", "active").
    OrderBy("created_at", "DESC").
    Get(&users)
```

---

## Example

```go
var users []User

err := database.Table("users").
    Where("status", "active").
    OrderBy("created_at", "DESC").
    Get(&users)

if err != nil {
    return err
}
```
