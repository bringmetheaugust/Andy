package main

import (
	"andy/src/keys"
	"os"
	"strings"
)

func main() {
	// defer byeBye()

	args := os.Args[1:]
	var argsP *[]string = &args

	if len(*argsP) == 0 {
		greetings()
		os.Exit(0)
	}

	switch true {
	case strings.HasSuffix((*argsP)[0], "--help") || strings.HasSuffix((*argsP)[0], "-h"):
		keys.Help()
	case strings.HasSuffix((*argsP)[0], "--version") || strings.HasSuffix((*argsP)[0], "-v"):
		keys.Version()
	}
}

// func byeBye() {
// 	fmt.Println("Bye bye!")
// }
