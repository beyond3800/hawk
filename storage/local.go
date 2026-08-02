package storage

import (
	"io"
	"os"
	"path/filepath"
)

type LocalDriver struct {
	root string
}

func Local(root string) *LocalDriver {
	return &LocalDriver{
		root: root,
	}
}

func (l *LocalDriver) Put(path string, file io.Reader) error {

	fullPath := filepath.Join(l.root, path)

	dir := filepath.Dir(fullPath)

	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	dst, err := os.Create(fullPath)
	if err != nil {
		return err
	}

	defer dst.Close()

	_, err = io.Copy(dst, file)

	return err
}
func (l *LocalDriver) Get(path string) (io.ReadCloser, error) {
	fullPath := filepath.Join(l.root, path)

	file, err := os.Open(fullPath)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (l *LocalDriver) Delete(path string) error {
	fullPath := filepath.Join(l.root, path)

	return os.Remove(fullPath)
}

func (l *LocalDriver) Exists(path string) bool {
	fullPath := filepath.Join(l.root, path)

	_, err := os.Stat(fullPath)

	return err == nil
}

func (l *LocalDriver) URL(path string) string {
	return "/" + filepath.ToSlash(filepath.Join(l.root, path))
}