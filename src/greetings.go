package main

import (
	"andy/src/helper"
	"fmt"
	"io/ioutil"
)

func greetings() {
	helpTxt, err := ioutil.ReadFile("./static/greetings.txt")

	if err != nil {
		helper.PanicMessage()
	}

	fmt.Println(string(helpTxt))
}
