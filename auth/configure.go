package auth

import "errors"

var authConfig Config

func Configure(config Config) error {
	if config.SecretKey == "" {
		return errors.New("auth: secret key cannot be empty")
	}
	authConfig = config

	return nil
}