package seabattle

import (
	"andy/pkg/colorPrint"
	"strconv"
)

type game struct {
	net    net   // Array of cells
	steps  uint8 // Step's count
	isOver bool  // Game is over
}

func (g *game) checkIsOver() {
	if g.steps <= 0 {
		g.isOver = true
	}
}

// Get info about next step from CLI
func (g game) requestStep() step {
	var newStep step

	collumn := newStep.requestStepCollumn()
	raw := newStep.requestStepRaw()

	return step{collumn, raw}
}

// Request next step, update net, check main game state.
// Return step result: {true} if step is successfull
func (g *game) makeStep() bool {
	colorPrint.GreenPrint("You have only " + strconv.Itoa(int(g.steps)) + " steps.\n")
	step := g.requestStep()

	currentCell := &(g.net[step.collumn][step.raw])

	if currentCell.isChecked {
		colorPrint.YellowPrint("Current cell is used. Please, try one more time.\n")

		return false
	} else {
		(*currentCell).isChecked = true
		g.net.build()
		g.steps = g.steps - 1
		g.checkIsOver()

		return true
	}
}
