package options

import (
	"fmt"
	"io/ioutil"
)

// Help show help (description, list of arguments, options and flags)
func Help() {
	helpTxt, err := ioutil.ReadFile("./static/help.txt")

	if err != nil {
		fmt.Println("pizda..")
	}

	fmt.Println(string(helpTxt))
}
