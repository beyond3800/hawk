package types

import (
	"time"
)

type Token struct {
	User_id string    `json:"user_id"`
	Jti     string    `json:"jti"`
	Exp     time.Time `json:"exp"`
}
