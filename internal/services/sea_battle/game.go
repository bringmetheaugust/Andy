package seabattle

import (
	"andy/pkg/colorPrint"
	"fmt"
	"strconv"
)

type game struct {
	net             net     // Array of cells
	ships           shipMap // Map of ships
	steps           uint8   // Step's count
	isOver          bool    // Game is over
	win             bool    // User game result
	shipsAreVisible bool    // Show ships (for dev mode or after finishing battle)
}

// Check if all ships are destroyed
func (g *game) isAllShipsAreDestroyed() bool {
	var destroyedShipACount int

	for _, ship := range g.ships {
		if ship.isDestroyed {
			destroyedShipACount++
		}
	}

	if len(g.ships) == destroyedShipACount {
		return true
	} else {
		return false
	}
}

// Check steps amount, all ships statuses
func (g *game) checkGameProcess() {
	if g.isAllShipsAreDestroyed() {
		g.win = true
	}

	if g.steps <= 0 || g.win {
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

// Request steps amount from CLI
func (g *game) requestStepAmount() {
	var stepCountInput string

	fmt.Print("Max step's value is " + strconv.Itoa(int(maxStepCount)) + ". How many steps do U need?: ")
	fmt.Scanln(&stepCountInput)

	stepCount, err := strconv.Atoi(stepCountInput)

	if err != nil || stepCount <= 0 || uint8(stepCount) > maxStepCount {
		panic("Plaese, enter correct steps number.")
	}

	g.steps = uint8(stepCount)
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
		currentCell.isChecked = true

		if currentCell.hasShip {
			g.checkShip(currentCell)
		}

		g.steps = g.steps - 1
		g.net.print(g.shipsAreVisible)
		g.checkGameProcess()

		return true
	}
}

// map ship status and ship's coordinates
func (g *game) checkShip(selectedCell *cell) {
	currentShip, ok := g.ships[selectedCell.shipId]

	if !ok {
		panic("Ship not found.")
	}

	var isDestroyed uint8
	shipCoordinates := currentShip.coordinates

	for _, c := range shipCoordinates {
		if g.net[c.collumn][c.row].isChecked {
			isDestroyed++
		}
	}

	// check if all ship's cell are selected and mark them as destroyed
	if len(shipCoordinates) == int(isDestroyed) {
		for _, c := range shipCoordinates {
			g.net[c.collumn][c.row].shipStatus = 2
			currentShip.isDestroyed = true
		}
	} else {
		selectedCell.shipStatus = 1
	}
}

// Generate ships and inject them to the net
func (g *game) genShips() {
	// map ship list data
	for shipLen, shipAmount := range shipsConfig {

		// map each ship by amount
		for i := uint8(0); i < shipAmount; i++ {
			var tryCount uint8
			var newShip ship

			// generate ships while it's coordinates are free
			for {
				if tryCount > 100 {
					panic("To many tries to generate ships. Sorry. Try one more time.")
				}

				var isExisted bool
				newShip = ship{}.New(shipLen)

				for _, c := range newShip.coordinates {
					if g.net[c.collumn][c.row].hasShip {
						isExisted = true
						break
					}
				}

				if !isExisted {
					break
				} else {
					tryCount++
				}
			}

			g.ships[newShip.id] = &newShip

			// map ship's coordinates on net
			for _, coordinate := range newShip.coordinates {
				g.net[coordinate.collumn][coordinate.row] = cell{
					hasShip:    true,
					shipStatus: 0,
					shipId:     newShip.id,
				}
			}
		}
	}
}
