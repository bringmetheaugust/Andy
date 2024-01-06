package seabattle

import (
	"andy/pkg/colorPrint"
	"fmt"
	"strconv"
	"strings"
)

type step struct {
	collumn int
	raw     int
}

func (g step) requestStepCollumn() int {
	var input string

	fmt.Print("Enter collumn(from 0 to 9): ")
	fmt.Scanln(&input)

	collumn, _ := strconv.Atoi(input)

	// todo: input errors
	if collumn < 0 || collumn > netGrid-1 {
		colorPrint.RedPrint("Collumn number is out of diapazone. Please, input correct collumn number!\n")
	}

	fmt.Println(collumn)

	return collumn
}

func (g step) requestStepRaw() int {
	var input string

	fmt.Printf("Enter raw (from %v to %v)): ", collumnChars[0], collumnChars[9])
	fmt.Scanln(&input)

	rawIndex := -1

	for i, v := range collumnChars {
		if v == input || v == strings.ToUpper(input) {
			rawIndex = i
		}
	}

	// todo: input errors
	if rawIndex == -1 {
		colorPrint.RedPrint("Raw number is out of diapazone. Please, input correct raw number!\n")
	}

	return rawIndex
}
