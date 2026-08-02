package storage

type Config struct {
	Driver string
	Root   string
	BaseURL string
}

var defaultStorage *Storage

func Configure(cfg Config) error {

	switch cfg.Driver {

	case "local":
		defaultStorage = New(Local(cfg.Root))

	default:
		return ErrInvalidDriver
	}

	return nil
}

func Default() *Storage {
	return defaultStorage
}