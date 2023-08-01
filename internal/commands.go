package internal

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/amirhnajafiz/captain-mustache/pkg/filesystem"
	"github.com/amirhnajafiz/captain-mustache/pkg/logger"

	"github.com/spf13/cobra"
)

const (
	DockerFileAddress    = "src/runtime/Dockerfile"
	DockerComposeAddress = "src/runtime/docker-compose.yaml"
)

var (
	stubs = map[string]string{
		"redis":      "src/stubs/database/redis.yaml",
		"mongodb":    "src/stubs/database/mongodb.yaml",
		"postgresql": "src/stubs/database/postgresql.yaml",
		"mysql":      "src/stubs/database/mysql.yaml",
	}
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
			if ok, err := exists("build"); err != nil && !ok {
				r.Logger.Error(ErrNoBuildFile)

				return
			}

			root := exec.Command("docker", "compose", "up", "-d")

			root.Stdout = os.Stdout
			root.Stderr = os.Stderr

			if err := root.Start(); err != nil {
				r.Logger.Error(ErrSystemDocker)
			}

			time.Sleep(1 * time.Second)
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
				r.Logger.Error(ErrSystemDocker)
			}

			time.Sleep(1 * time.Second)
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
				r.Logger.Error(ErrSystemDocker)
			}

			time.Sleep(1 * time.Second)
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
			if ok, err := exists("main.go"); err != nil && !ok {
				r.Logger.Error(ErrMainFile)

				return
			}

			if ok, err := exists("go.mod"); err != nil && !ok {
				r.Logger.Error(ErrDirectoryStructure)

				return
			}

			if ok, err := exists("build"); err == nil && ok {
				r.Logger.Error(ErrDuplicateDirectory)

				return
			}

			if ok, err := exists("docker-compose.yaml"); err == nil && ok {
				r.Logger.Error(ErrDuplicateDirectory)

				return
			}

			if err := os.Mkdir("build", 0750); err != nil {
				r.Logger.Error(ErrDuplicateDirectory)

				return
			}

			command, err := getInputs()
			if err != nil {
				r.Logger.Error(err)

				return
			}

			dockerFile, err := filesystem.ReadFile(DockerFileAddress)
			if err != nil {
				r.Logger.Error(ErrModuleNotFound)

				return
			}

			// golang version, os, arch
			dockerFile = strings.Replace(dockerFile, "{{version}}", command.GoVersion, 1)
			dockerFile = strings.Replace(dockerFile, "{{GOOS}}", command.OperatingSystem, 1)
			dockerFile = strings.Replace(dockerFile, "{{GOARCH}}", command.Architecture, 1)

			if er := filesystem.WriteFile(dockerFile, "build/Dockerfile"); er != nil {
				r.Logger.Error(er)

				return
			}

			dockerComposeFile, err := filesystem.ReadFile(DockerComposeAddress)
			if err != nil {
				r.Logger.Error(ErrModuleNotFound)

				return
			}

			dockerComposeFile = strings.Replace(dockerComposeFile, "{{port}}", command.Port, 2)

			str := ""

			for _, c := range command.SubCommands {
				tmp, er := filesystem.ReadFile(stubs[c])
				if er != nil {
					r.Logger.Error(er)

					return
				}

				str = fmt.Sprintf("%s\n%s", str, tmp)
			}

			dockerComposeFile = strings.Replace(dockerComposeFile, "{{stubs}}", str, 1)

			if er := filesystem.WriteFile(dockerComposeFile, "docker-compose.yaml"); er != nil {
				r.Logger.Error(er)

				return
			}

			r.Logger.Info("dockerfiles created!")
		},
	}
}
