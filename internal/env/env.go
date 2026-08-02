package env

import (
	"fmt"
)

var current *Parser

func Load(path string) error {
	e := New(path)
	if err := e.Load(); err != nil {
		return err
	}
	current = e
	return nil
}

func Get(key string) (string, bool) {
	if current == nil {
		return "", false
	}
	return current.Get(key)
}

func Has(key string) bool {
	if current == nil {
		return false
	}
	return current.Has(key)
}

func Set(key, value string) {

	if current == nil {
		return
	}
	current.Set(key, value)
}

func Save() error {
	if current == nil {
		return fmt.Errorf("environment not loaded")
	}
	return current.Save()
}

func Bool(key string) (bool, error) {
	value, ok := Get(key)
	if !ok {
		return false, fmt.Errorf("key %s not found", key)
	}
	if value == "true" || value == "1" {
		return true, nil
	}
	if value == "false" || value == "0" {
		return false, nil
	}
	return false, fmt.Errorf("invalid boolean value for key %s: %s", key, value)
}

