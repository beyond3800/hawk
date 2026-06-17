
---

# `docs/validation.md`

```md
# Validation

## Validate Request Data

```go
Rules allowed
required|email|min:num|max:num|unique:table
type Users struct {
    ID          string        `json:"id" db:"id"`
    Name        string        `json:"name" validate:"required" db:"name"`
    Email       string        `json:"email" validate:"required|email" db:"email"`
    Password    string        `json:"password" validate:"required|min:6" db:"password"`
    Created_at  time.Time     `json:"created_at"`
    Updated_at  sql.NullTime  `json:"updated_at"`
}

var user Users
	
	if err := c.BindAndValidate(&user); err != nil {
		c.ValidationError(err)
		return
	}