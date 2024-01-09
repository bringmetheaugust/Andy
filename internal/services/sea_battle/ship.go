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
	1: 1,
	// 2: 3,
	// 3: 2,
	// 4: 1,
}

// Generate new ship
// TODO: check if cell is зайнята
// TODO: build 2/3/4 length ships
func (s ship) New(shipLen uint8) ship {
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
		id:          shipId(tools.GenId()),
		coordinates: cellCoordinates,
	}
}
