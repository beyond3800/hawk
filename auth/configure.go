package auth

import "errors"

var cgf Config

func Configure(config Config) error {
	if config.SecretKey == "" {
		return errors.New("auth: secret key cannot be empty")
	}
	cgf = config

	return nil
}