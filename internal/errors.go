package internal

import "errors"

var (
	ErrDirectoryStructue = errors.New("no go project in current directory")
	ErrMainFile          = errors.New("failed to find main.go")
	ErrNoBuildFile       = errors.New("project is not build yet")
	ErrSystemDocker      = errors.New("your system does not have docker demon on")
	ErrWrongInput        = errors.New("command did not match any")
)
