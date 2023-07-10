package filesystem

import (
	"bufio"
	"io"
	"os"
)

func ReadFile(path string) (io.Reader, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, ErrFileNotFound
	}

	return bufio.NewReader(f), nil
}

func WriteFile(reader io.Reader, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return ErrCreateFile
	}

	w := bufio.NewWriter(f)

	if _, err := w.ReadFrom(reader); err != nil {
		return ErrSystemFile
	}

	return nil
}
