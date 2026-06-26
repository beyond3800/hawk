# Query Builder

The Hawk Query Builder provides a fluent interface for interacting with your database.

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
database.Table("users").Insert(map[string]any{
    "name":  "Adam",
    "email": "adam@example.com",
})
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
