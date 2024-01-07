package seabattle

import (
	"andy/pkg/colorPrint"
	"fmt"
	"strconv"
	"strings"
)

type step struct {
	collumn uint8 // Selected collumn index
	raw     uint8 // Selected raw index
}

// Request and parse collumn index from CLI
func (g step) requestStepCollumn() uint8 {
	var input string

	fmt.Print("Enter collumn(from 0 to 9): ")
	fmt.Scanln(&input)

	collumnInt, _ := strconv.Atoi(input)
	collumn := uint8(collumnInt)

	if collumn < 0 || collumn > netGrid-1 {
		colorPrint.RedPrint("Collumn number is out of diapazone. Please, input correct collumn number!\n")
		panic("TODO")
	}

	return collumn
}

// Request and parse raw index from CLI
func (g step) requestStepRaw() uint8 {
	var input string

	fmt.Printf("Enter raw (from %v to %v)): ", collumnChars[0], collumnChars[9])
	fmt.Scanln(&input)

	var rawIndex int8 = -1

	for i, v := range collumnChars {
		if v == input || v == strings.ToUpper(input) {
			rawIndex = int8(i)
		}
	}

	if rawIndex == -1 {
		colorPrint.RedPrint("Raw number is out of diapazone. Please, input correct raw number!\n")
		panic("TODO")
	}

	return uint8(rawIndex)
}
