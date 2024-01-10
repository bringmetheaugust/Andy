package seabattle

type cell struct {
	isChecked  bool // Cell is used
	hasShip    bool // Has a ship
	shipId          // shipID
	shipStatus      // Ship status
}

// Coordinate type
type coordinate uint8

type coordinates struct {
	row     coordinate // Row array index
	collumn coordinate // Collumn array index
}
