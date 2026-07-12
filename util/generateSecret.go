package util

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateSecret() (string, error) {
	key := make([]byte, 32)

	if _, err := rand.Read(key); err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(key), nil
}