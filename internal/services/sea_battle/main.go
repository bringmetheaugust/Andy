package seabattle

import (
	"fmt"
)

const maxStepCount uint8 = 100

func StartGame(withDev bool) {
	g := game{
		ships:           make(shipMap),
		shipsAreVisible: withDev,
	}

	g.requestStepAmount()
	g.genShips()
	g.net.print(g.shipsAreVisible)

	for {
		if g.isOver {
			if !g.win {
				fmt.Println("You lose. Here is a map of a completed battle:")
				g.net.print(true)
			} else {
				fmt.Println("You win!!")
			}

			break
		}

		for {
			if g.makeStep() {
				break
			}
		}
	}
}
