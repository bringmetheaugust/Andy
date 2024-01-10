package seabattle

import (
	"fmt"
)

const maxStepCount uint8 = 100

func StartGame(withDev bool) {
	game := game{
		ships:           make(shipMap),
		shipsAreVisible: withDev,
	}

	game.requestStepAmount()
	game.genShips()
	game.net.print(game.shipsAreVisible)

	for {
		if game.isOver {
			if !game.win {
				fmt.Println("You lose. Here is a map of a completed battle:")
				game.net.print(true)
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
