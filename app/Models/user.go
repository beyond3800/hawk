package Models

import (
	"database/sql"
	"time"
)

type User struct {
	ID         string `json:"id" validate:"required" db:"id"`
	Name       string `json:"firstName" validate:"required" db:"name"`
	Email      string `json:"email" validate:"required|email" db:"email"`
	Password   string `json:"password" validate:"required|min:6" db:"password"`
	Created_at time.Time `json:"created_at" db:"created_at"`
	Updated_at sql.NullTime `json:"updated_at" db:"updated_at"`
}

