package auth

import "os"

func Init(config Config) {
	os.Setenv("APP_SECRET_KEY", config.SecretKey)
	os.Setenv("APP_ISSUER", config.Issuer)
}