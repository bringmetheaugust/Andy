package seabattle

import "andy/pkg/colorPrint"

func StartGame() {
	game := game{
		steps: 10,
	}

	game.net.build()

	for {
		if game.isOver {
			colorPrint.BluePrint("You lose...\n")
			break
		}

		for {
			if game.makeStep() {
				break
			}
		}
	}
}
