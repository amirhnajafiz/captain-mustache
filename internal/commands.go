package internal

import (
	"os"
	"os/exec"

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
			// todo: implement the base logic
		},
	}
}
