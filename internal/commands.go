package internal

import (
	"os/exec"

	"github.com/spf13/cobra"
)

// UpCommand starts docker containers.
func UpCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "up",
		Short: "start containers",
		Long:  "start docker containers by using docker compose up",
		Run: func(cmd *cobra.Command, args []string) {
			root := exec.Command("docker", "compose", "up", "-d")

			if err := root.Start(); err != nil {
				// todo: error log
			}
		},
	}
}

// DownCommand stops docker containers.
func DownCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "down",
		Short: "stop containers",
		Long:  "stop docker containers by using docker compose down",
		Run: func(cmd *cobra.Command, args []string) {
			root := exec.Command("docker", "compose", "down")

			if err := root.Start(); err != nil {
				// todo: error log
			}
		},
	}
}

// StatusCommand returns a status of containers.
func StatusCommand() {

}

// BuildCommand generates docker compose manifest.
func BuildCommand() {

}
