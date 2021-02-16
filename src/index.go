package main

import (
	"fmt"
	"os"
)

func main() {
	defer byeBye()

	args := os.Args
	// var argsP *[]string = &args

	if len(args) == 1 {
		greetings()
		os.Exit(0)
	}
}

func byeBye() {
	fmt.Println("Bye Bye🛌🏻...")
}
