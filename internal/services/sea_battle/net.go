package seabattle

import (
	"andy/pkg/colorPrint"
	"fmt"
	"strconv"
	"strings"
)

const netGrid uint8 = 10
const cellSymbol string = "* "

type net [netGrid][netGrid]cell

var alphabet = [...]string{
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
	"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
}
var collumnChars = alphabet[:netGrid]

// Print net to CLI from game state.
// {isHidden} parameter show all hidden ships (for game over or testing view)
func (n net) build(isHidden bool) {
	colorPrint.BluePrint("   " + strings.Join(collumnChars[:], " ") + "\n")

	for i, coll := range n {
		colorPrint.BluePrint(strconv.Itoa(i) + "| ")

		for _, v := range coll {
			switch {
			case v.isChecked && v.hasShip && v.ship.isDestroyed: // destroyed ship
				colorPrint.RedPrint(cellSymbol)
			case v.isChecked && v.hasShip: // injured ship
				colorPrint.YellowPrint(cellSymbol)
			case isHidden && v.isChecked: // used empty cell
				colorPrint.BluePrint(cellSymbol)
			case !isHidden && v.hasShip && !v.isDestroyed: // surviving part of the ship
				colorPrint.GreenPrint(cellSymbol)
			default: //none selected cell
				fmt.Printf(cellSymbol)
			}
		}

		fmt.Printf("\n")
	}
}
