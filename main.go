package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
)

func main() {
	var (
		CommandFlag = flag.String("command", "build", "mustache command (build, up, down)")
	)

	flag.Parse()

	command := exec.Command("./bin/mustache", *CommandFlag)

	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	if er := command.Start(); er != nil {
		log.Fatal(er)
	}
}
