package filesystem

import "errors"

var (
	ErrFileNotFound = errors.New("file not found")
	ErrCreateFile   = errors.New("failed to create file")
	ErrSystemFile   = errors.New("failed to write into file")
)
