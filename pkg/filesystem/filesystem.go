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
