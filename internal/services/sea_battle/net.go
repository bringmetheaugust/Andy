package seabattle

import (
	"andy/pkg/colorPrint"
	"fmt"
	"strconv"
	"strings"
)

const netGrid int = 10
const cellSymbol string = "* "

type net [netGrid][netGrid]cell

var alphabet = [...]string{
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
	"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
}
var collumnChars = alphabet[:netGrid]

func (n net) build() {
	colorPrint.BluePrint("   " + strings.Join(collumnChars[:], " ") + "\n")

	for i, coll := range n {
		colorPrint.BluePrint(strconv.Itoa(i) + "| ")

		for _, v := range coll {
			if v.isChecked {
				colorPrint.RedPrint(cellSymbol)
			} else {
				fmt.Printf(cellSymbol)
			}
		}

		fmt.Printf("\n")
	}
}
