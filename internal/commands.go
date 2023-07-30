package internal

import (
	"github.com/amirhnajafiz/captain-mustache/pkg/filesystem"
	"os"
	"os/exec"
	"strings"

	"github.com/amirhnajafiz/captain-mustache/pkg/logger"

	"github.com/spf13/cobra"
)

type Root struct {
	Logger logger.Logger
}

// UpCommand starts docker containers.
func (r Root) UpCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "up",
		Short: "start containers",
		Long:  "start docker containers by using docker compose up",
		Run: func(cmd *cobra.Command, args []string) {
			root := exec.Command("docker", "compose", "up", "-d")

			root.Stdout = os.Stdout
			root.Stderr = os.Stderr

			if err := root.Start(); err != nil {
				r.Logger.Error(err)
			}
		},
	}
}

// DownCommand stops docker containers.
func (r Root) DownCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "down",
		Short: "stop containers",
		Long:  "stop docker containers by using docker compose down",
		Run: func(cmd *cobra.Command, args []string) {
			root := exec.Command("docker", "compose", "down")

			root.Stdout = os.Stdout
			root.Stderr = os.Stderr

			if err := root.Start(); err != nil {
				r.Logger.Error(err)
			}
		},
	}
}

// StatusCommand returns a status of containers.
func (r Root) StatusCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "status",
		Short: "containers status",
		Long:  "display docker containers status by using docker ps",
		Run: func(cmd *cobra.Command, args []string) {
			root := exec.Command("docker", "ps", "-a")

			root.Stdout = os.Stdout
			root.Stderr = os.Stderr

			if err := root.Start(); err != nil {
				r.Logger.Error(err)
			}
		},
	}
}

// BuildCommand generates docker compose manifest.
func (r Root) BuildCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "build",
		Short: "create manifest",
		Long:  "create docker compose and dockerfile",
		Run: func(cmd *cobra.Command, args []string) {
			if ok, err := exists("build"); err != nil && !ok {
				r.Logger.Error(ErrDuplicateDirectory)

				return
			}

			if err := os.Mkdir("build", 0750); err != nil {
				panic(err)
			}

			command, err := importBaseQuestions()
			if err != nil {
				panic(err)
			}

			f, err := filesystem.ReadFile("src/runtime/Dockerfile")
			if err != nil {
				r.Logger.Error(ErrModuleNotFound)

				return
			}

			// golang version, os, arch
			f = strings.Replace(f, "{{version}}", command.Imports.GoVersion, 1)
			f = strings.Replace(f, "{{GOOS}}", command.Imports.OperatingSystem, 1)
			f = strings.Replace(f, "{{GOARCH}}", command.Imports.Architecture, 1)

			if er := filesystem.WriteFile(f, "build/Dockerfile"); er != nil {
				panic(err)
			}

			f, err = filesystem.ReadFile("src/runtime/docker-compose.yaml")
			if err != nil {
				r.Logger.Error(ErrModuleNotFound)

				return
			}

			// todo: process files with inputs

			if err := filesystem.WriteFile("docker-compose.yaml", f); err != nil {
				panic(err)
			}
		},
	}
}
