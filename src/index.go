package main

import (
	"andy/src/commands"
	"andy/src/keys"
	"andy/src/utils"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		utils.Greetings()
		os.Exit(0)
	}

	firstArg := args[0]
	var argsP *[]string = &args

	switch true {
	case firstArg == "--help" || firstArg == "-h":
		keys.Help()
	case firstArg == "--version" || firstArg == "-v":
		keys.Version()
	case firstArg == "server":
		commands.Server.Run((*argsP)[1:])
	}
}
