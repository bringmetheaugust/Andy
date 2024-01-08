package seabattle

type cell struct {
	isChecked bool // Cell is used
	hasShip   bool // Has a ship
	ship           // Ship data
}

type coordinates struct {
	row     uint8
	collumn uint8
}

type ship struct {
	id          int64
	isDestroyed bool          // Ship is destroyed
	coordinates []coordinates // Coordinates of all ship's cells
}
