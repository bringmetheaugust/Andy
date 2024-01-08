package seabattle

import (
	"fmt"
	"strconv"
)

const maxStepCount uint8 = 100

func StartGame() {
	var stepCountInput string

	fmt.Print("Max step's value is " + strconv.Itoa(int(maxStepCount)) + ". How many steps do U need?: ")
	fmt.Scanln(&stepCountInput)

	stepCount, err := strconv.Atoi(stepCountInput)

	if err != nil || stepCount <= 0 || uint8(stepCount) > maxStepCount {
		panic("Plaese, enter correct steps number.")
	}

	game := game{
		steps: uint8(stepCount),
	}

	game.genShips()
	game.net.build(true)

	for {
		if game.isOver {
			if !game.win {
				fmt.Println("You lose. Here is a map of a completed battle:")
				game.net.build(false)
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
