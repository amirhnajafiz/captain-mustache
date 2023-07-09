package main

import "flag"

func main() {
	var (
		CommandFlag = flag.String("command", "build", "mustache command (build, up, down)")
	)

	flag.Parse()
}
