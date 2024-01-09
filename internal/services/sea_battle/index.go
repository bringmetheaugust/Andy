package seabattle

import (
	"fmt"
)

const maxStepCount uint8 = 100

func StartGame() {
	game := game{
		ships: make(shipMap),
	}

	game.requestStepAmount()
	game.genShips()
	game.net.print(true)

	for {
		if game.isOver {
			if !game.win {
				fmt.Println("You lose. Here is a map of a completed battle:")
				game.net.print(false)
			} else {
				fmt.Println("You win!!")
			}

			break
		}

		for {
			if game.makeStep() {
				break
			}
		}
	}
}
