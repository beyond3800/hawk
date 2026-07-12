package auth

import (
	"fmt"
	"time"

	"github.com/beyond3800/hawk/id"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
    UserID string `json:"user_id"`
    jwt.RegisteredClaims
}

func CreateToken(userID string, expiresIn time.Duration) (string, error) {
	if authConfig.SecretKey == "" {
		return "", fmt.Errorf("auth package has not been initialized")
	}

	now := time.Now()

	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        id.New(),
			Subject:   userID,
			Issuer:    authConfig.Issuer,
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(expiresIn)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(authConfig.SecretKey))
}