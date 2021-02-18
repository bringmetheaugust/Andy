package keys

import (
	"andy/src/helper"
	"fmt"
	"io/ioutil"
)

// Help show help (description, list of arguments, options and flags)
func Help() {
	helpTxt, err := ioutil.ReadFile("./static/help.txt")

	if err != nil {
		helper.PanicMessage()
	}

	fmt.Println(string(helpTxt))
}
