package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

type command struct {
	title   string
	command string
}

func main() {
	args := os.Args

	if args[1] == "help" {
		helpTxt, err := ioutil.ReadFile("./docs/help.txt")

		if err != nil {
			fmt.Println("pizda..")
		}

		fmt.Println(string(helpTxt))
	}

	os.Exit(0)
}
