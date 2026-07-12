package types

import (
	"time"

	"github.com/beyond3800/hawk/auth"
)

type Token struct {
	User_id string    `json:"user_id"`
	Jti     string    `json:"jti"`
	Exp     time.Time `json:"exp"`
}
type EnvConfig struct{
	Auth auth.Config
} 