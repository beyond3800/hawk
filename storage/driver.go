package storage

import "io"

type Driver interface {
	Put(path string, file io.Reader) error
	Get(path string) (io.ReadCloser, error)
	Delete(path string) error
	Exists(path string) bool
	URL(path string) string
}