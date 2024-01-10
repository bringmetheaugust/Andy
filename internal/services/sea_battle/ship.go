package seabattle

import (
	"andy/tools"
	"math/rand"
)

type shipId int64

// 0 - целый, 1 - подбомбленый, 2 - убит
type shipStatus uint8

// Map of ships: { shipId: ship }
type shipMap map[shipId]*ship

type ship struct {
	id          shipId        // ship ID
	isDestroyed bool          // Ship is destroyed
	coordinates []coordinates // Coordinates of all ship's cells
}

// Map of ships. Key is a ship cell lenght, value is a count of this ship's type
var shipsConfig = map[uint8]uint8{
	1: 5,
	// 2: 3,
	// 3: 2,
	// 4: 1,
}

// Generate new ship with random coordinates
// TODO: generate long ships
func (s ship) New(shipLen uint8) ship {
	var cellCoordinates []coordinates
	var randRowIndex = rand.Intn(int(netGrid))     // entry point for ship's coordinates
	var randCollumnIndex = rand.Intn(int(netGrid)) // entry point for ship's coordinates

	for i := uint8(0); i < shipLen; i++ {
		cellCoordinates = append(cellCoordinates, coordinates{
			row:     coordinate(randRowIndex),
			collumn: coordinate(randCollumnIndex),
		})
	}

	return ship{
		id:          shipId(tools.GenId()),
		coordinates: cellCoordinates,
	}
}
