package main

import (
	"github.com/amirhnajafiz/captain-mustache/internal"
	"github.com/amirhnajafiz/captain-mustache/pkg/logger"

	"github.com/spf13/cobra"
)

func main() {
	// init logger
	l := logger.Logger{
		Mode: logger.DebugMode,
	}

	// init root commands
	root := internal.Root{
		Logger: l,
	}

	// create cobra root
	cmd := cobra.Command{}

	cmd.AddCommand(
		root.DownCommand(),
		root.UpCommand(),
		root.StatusCommand(),
		root.BuildCommand(),
	)

	// execute commands
	if err := cmd.Execute(); err != nil {
		l.Error(err)
	}
}
