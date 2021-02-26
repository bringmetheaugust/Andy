package keys

import (
	"andy/src/constants"
	"fmt"
)

// Help show help (description, list of arguments, options and flags)
func Help() {
	fmt.Println(constants.Help)
}
