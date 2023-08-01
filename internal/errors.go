package internal

import "errors"

var (
	ErrDirectoryStructure = errors.New("no go project in current directory")
	ErrMainFile           = errors.New("failed to find main.go")
	ErrNoBuildFile        = errors.New("project is not build yet")
	ErrSystemDocker       = errors.New("your system does not have docker demon on")
	ErrDuplicateDirectory = errors.New("build directory already exists")
	ErrModuleNotFound     = errors.New("strago modules not found")
)
