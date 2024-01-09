package seabattle

type cell struct {
	isChecked  bool // Cell is used
	hasShip    bool // Has a ship
	shipId          // shipID
	shipStatus      // Ship status
}

type coordinates struct {
	row     uint8 // Row array index
	collumn uint8 // Collumn array index
}
