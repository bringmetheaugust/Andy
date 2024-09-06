package snake

import "math/rand"

func getRandomPoint(max int) (x, y int) {
	x = rand.Intn(max)
	y = rand.Intn(max)

	return
}
