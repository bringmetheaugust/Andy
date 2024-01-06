package seabattle

func StartGame() {
	game := game{}

	game.net.build()
	game.makeStep()
}
