
---

# `docs/query-builder.md`

```md
# Query Builder

## Select

```go
var users []User

err := database.Table("users").
    Get(&users)

var user User

err := database.Table("users").
    Where("id", 1).
    First(&user)
// orW
database.Table("users").
    Where("name", "John").
    OrWhere("name", "Jane")

database.Table("users").
    OrderBy("created_at", "DESC")

database.Table("users").Insert(map[string]any{
    "name": "Adam",
    "email": "adam@example.com",
})

database.Table("users").
    Where("id", 1).
    Update(map[string]any{
        "name": "Updated Name",
    })

database.Table("users").
    Where("id", 1).
    Delete()