package filesystem

import (
	"bufio"
	"io"
	"os"
)

func ReadFile(path string) (string, error) {
	f, err := os.ReadFile(path)
	if err != nil {
		return "", ErrFileNotFound
	}

	return string(f), nil
}

func WriteFile(reader io.Reader, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return ErrCreateFile
	}

	w := bufio.NewWriter(f)

	if _, er := w.ReadFrom(reader); er != nil {
		return ErrSystemFile
	}

	return nil
}
