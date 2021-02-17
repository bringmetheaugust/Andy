package main

import (
	"os"
)

type arg struct {
}

func main() {
	// defer byeBye()

	args := os.Args[1:]
	var argsP *[]string = &args

	switch true {
	case len(*argsP) == 0:
		greetings()
		os.Exit(0)
		// case args[1] == :
		// 	fmt.Println("Linux.")
		// default:
		// 	// freebsd, openbsd,
		// 	// plan9, windows...
		// 	fmt.Printf("%s.\n", os)
	}
}

// func byeBye() {
// 	fmt.Println("Bye bye!")
// }
