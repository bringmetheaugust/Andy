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

// Print net to CLI from game state. {shipsAreVisible} parameter show all hidden ships (for game over or testing view)
func (n net) print(shipsAreVisible bool) {
	colorPrint.BluePrint("   " + strings.Join(collumnChars[:], " ") + "\n")

	for i, coll := range n {
		colorPrint.BluePrint(strconv.Itoa(i) + "| ")

		for _, cell := range coll {
			switch {
			case cell.hasShip:
				switch cell.shipStatus {
				case 2: // destroyed ship
					colorPrint.RedPrint(cellSymbol)
				case 1: // injured ship
					colorPrint.YellowPrint(cellSymbol)
				case 0: // surviving part of the ship
					if shipsAreVisible {
						colorPrint.GreenPrint(cellSymbol)
					} else {
						fmt.Printf(cellSymbol)
					}
				}
			case !cell.hasShip && cell.isChecked: // used empty cell
				colorPrint.BluePrint(cellSymbol)
			default: //none selected cell
				fmt.Printf(cellSymbol)
			}
		}

		fmt.Printf("\n")
	}
}
