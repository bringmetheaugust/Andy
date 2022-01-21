package commands

import (
	"fmt"
	"os"
)

func ParseCMD() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Hi! I'm Andy🐼!\nRun --help or -h to watch my commands.")
		os.Exit(0)
	}

	switch os.Args[1] {
	case "docker_server":
		RunDockerStatusServer()
	default:
		fmt.Println("unknown..")
	}
}
