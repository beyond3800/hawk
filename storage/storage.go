package storage

import "io"

type Storage struct {
	driver Driver
}

func New(driver Driver) *Storage {
	return &Storage{
		driver: driver,
	}
}

func (s *Storage) Put(path string, file io.Reader) error {
	return s.driver.Put(path, file)
}

func (s *Storage) Get(path string) (io.ReadCloser, error) {
	return s.driver.Get(path)
}

func (s *Storage) Delete(path string) error {
	return s.driver.Delete(path)
}

func (s *Storage) Exists(path string) bool {
	return s.driver.Exists(path)
}

func (s *Storage) URL(path string) string {
	return s.driver.URL(path)
}
