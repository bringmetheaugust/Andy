package snake

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type speed int

const (
	minGridSize = 3
	maxGridSize = 50
)

func requestGridSize() (g int) {
	fmt.Printf("Please, choose grid size (%d - %d): ", minGridSize, maxGridSize)

	for {
		s, err := readIntFromStd()

		if err != nil {
			fmt.Println("You have typed invalid character. Please, type number: ")
			continue
		}

		if s < minGridSize || s > maxGridSize {
			fmt.Printf("Please, type correct grid size (%d - %d): ", minGridSize, maxGridSize)
			continue
		}

		g = s

		break
	}

	return
}

func requestSpeed() (s speed) {
	fmt.Printf("Please, choose game speed (%d - low speed, %d - medium speed, %d - hight speed): ", lowSpeed, mediumSpeed, hightSpeed)

	for {
		p, err := readIntFromStd()

		if err != nil {
			fmt.Printf("You have typed invalid character. Please, type number (%d, %d or %d): ", hightSpeed, mediumSpeed, lowSpeed)
			continue
		}

		if p < int(hightSpeed) || p > int(lowSpeed) {
			fmt.Printf("Please, type correct speed (%d - %d): ", hightSpeed, lowSpeed)
			continue
		}

		s = speed(p)

		break
	}

	return
}

func readIntFromStd() (i int, err error) {
	r := bufio.NewReader(os.Stdin)
	l, _, _ := r.ReadLine()
	i, err = strconv.Atoi(string(l))

	return
}

func StartGame() {
	println("Use A/S/D/W keys or arrows for controls. Press q to quit game. And good luck to U!🐍")
	gridSize := requestGridSize()
	gridCenter := gridSize / 2
	speed := requestSpeed()
	gr := make(grid, gridSize)

	for r := 0; r < gridSize; r++ {
		gr[r] = make([]*cell, gridSize)

		for c := 0; c < gridSize; c++ {
			gr[r][c] = &cell{c, r}
		}
	}

	g := game{
		gridSize: gridSize,
		speed:    speed,
		grid:     gr,
	}
	g.snake = snakeCells{g.grid[gridCenter][gridCenter]}
	g.item = g.getRandomCell()

	g.run()
}
