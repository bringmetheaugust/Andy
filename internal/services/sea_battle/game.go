package seabattle

import (
	"andy/pkg/colorPrint"
	"andy/tools"
	"math/rand"
	"strconv"
)

type game struct {
	net    net   // Array of cells
	steps  uint8 // Step's count
	isOver bool  // Game is over
	win    bool
}

// map of ships. key is a ship cell lenght, value is a count of this ship's type
var ships = map[uint8]uint8{
	1: 5,
	// 2: 3,
	// 3: 2,
	// 4: 1,
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
		colorPrint.YellowPrint("Current cell is already used. Please, try one more time.\n")

		return false
	} else {
		if currentCell.hasShip {
			currentCell.ship.isDestroyed = true
		} else {

		}

		(*currentCell).isChecked = true
		g.net.build(true)
		g.steps = g.steps - 1

		if g.steps <= 0 {
			g.isOver = true
		}

		return true
	}
}

// Generate one ship
// TODO: check if cell is зайнята
func (g *game) genShip(shipLen uint8) ship {
	var cellCoordinates []coordinates

	for i := uint8(0); i < shipLen; i++ {
		var randRowIndex = rand.Intn(int(netGrid))
		var randCollumnIndex = rand.Intn(int(netGrid))

		cellCoordinates = append(cellCoordinates, coordinates{
			row:     uint8(randRowIndex),
			collumn: uint8(randCollumnIndex),
		})
	}

	return ship{
		id:          tools.GenId(),
		coordinates: cellCoordinates,
	}
}

// Generate ships and inject them to the net
func (g *game) genShips() {
	// map ship list data
	for shipLen, shipAmount := range ships {

		// map each ship by amount
		for i := uint8(0); i < shipAmount; i++ {
			newShip := g.genShip(shipLen)

			// map ship's coordinates on net
			for _, coordinate := range newShip.coordinates {
				g.net[coordinate.collumn][coordinate.row] = cell{
					hasShip: true,
					ship:    newShip,
				}
			}
		}
	}
}
