# Validation

Validation allows you to validate incoming request data using struct tags.

---

## Available Rules

* `required`
* `email`
* `min:num`
* `max:num`
* `unique:table`

---

## Defining Validation Rules

```go
type User struct {
    ID         string       `json:"id" db:"id"`
    Name       string       `json:"name" validate:"required" db:"name"`
    Email      string       `json:"email" validate:"required|email" db:"email"`
    Password   string       `json:"password" validate:"required|min:6" db:"password"`
    CreatedAt  time.Time    `json:"created_at"`
    UpdatedAt  sql.NullTime `json:"updated_at"`
}
```

---

## Validating Requests

```go
var user User

if err := c.BindAndValidate(&user); err != nil {
    c.ValidationError(err)
    return
}
```

`BindAndValidate`:

1. Parses the JSON request body.
2. Validates the struct fields.
3. Returns validation errors.

---

## Example Request

```json
{
    "name": "Adam",
    "email": "adam@example.com",
    "password": "secret123"
}
```

---

## Validation Errors

Example response:

```json
{
    "message": "validation failed",
    "errors": {
        "email": "email is required",
        "password": "minimum length is 6"
    }
}
```

---

## Rule Examples

### Required

```go
Name string `validate:"required"`
```

### Email

```go
Email string `validate:"email"`
```

### Minimum Length

```go
Password string `validate:"min:6"`
```

### Maximum Length

```go
Name string `validate:"max:255"`
```

### Unique Value

```go
Email string `validate:"unique:users"`
```

Checks that the value does not already exist in the specified table.
