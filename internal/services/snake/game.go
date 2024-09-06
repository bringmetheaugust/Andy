package snake

import (
	"fmt"
	"log"
	"math"
	"os"
	"os/exec"
	"slices"
	"strings"
	"time"

	"github.com/eiannone/keyboard"
)

type game struct {
	gridSize              int
	speed                 speed
	grid                  grid
	snake                 snakeCells
	snakeBackLastPosition *cell // for append last snake cell if snake gets item
	item                  *cell
	direction             direction
	gameOver              bool
}

type grid [][]*cell

type snakeCells []*cell

type cell struct {
	x int
	y int
}

type direction int

const (
	up direction = iota
	down
	left
	right
)

const (
	hightSpeed speed = iota + 1
	mediumSpeed
	lowSpeed
)

func (g *game) run() {
	defer func() {
		var sb strings.Builder
		sb.WriteRune('\n')

		if err := recover(); err != nil {
			sb.WriteString("You lose.. ")
		} else {
			sb.WriteString("You win! ")
		}

		sb.WriteString(fmt.Sprintf("Your score is %v points.", len(g.snake)))
		fmt.Println(sb.String())
		keyboard.Close()
	}()

	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	if err := keyboard.Open(); err != nil {
		log.Fatal(err)
	}

	go g.makeStep()

	for {
		if g.gameOver {
			return
		}

		g.checkStep()
		g.printGrid()
		<-time.After(time.Second * time.Duration(g.speed) / 3)
		g.move()
	}
}

func (g *game) printGrid() {
	var sb strings.Builder

	for _, rv := range g.grid {
		for _, cv := range rv {
			if cv == g.item {
				sb.WriteString("\033[1;35m *")
				continue
			}

			if slices.ContainsFunc(g.snake, func(c *cell) bool { return cv == c }) {
				sb.WriteString("\033[0;32m *")
			} else {
				sb.WriteString("\033[1;33m *")
			}
		}

		sb.WriteRune('\n')
	}

	fmt.Printf("\033[0;0H%v", sb.String())
}

func (g *game) move() {
	g.snakeBackLastPosition = g.snake[len(g.snake)-1]
	var prevPos cell

	for i, v := range g.snake {
		cPos := v

		if i == 0 {
			switch n := 0; g.direction {
			case up:
				n = (v.y - 1 + g.gridSize) % g.gridSize
				g.snake[i] = g.grid[n][v.x]
			case down:
				n = (v.y + 1) % g.gridSize
				g.snake[i] = g.grid[n][v.x]
			case left:
				n = (v.x - 1 + g.gridSize) % g.gridSize
				g.snake[i] = g.grid[v.y][n]
			case right:
				n = (v.x + 1) % g.gridSize
				g.snake[i] = g.grid[v.y][n]
			}
		} else {
			g.snake[i] = g.grid[prevPos.y][prevPos.x]
		}

		prevPos = *cPos
	}
}

func (g *game) makeStep() {
	defer func() {
		keyboard.Close()
		os.Exit(0)
	}()

	for {
		c, k, err := keyboard.GetKey()

		if err != nil {
			log.Fatal(err)
		}

		switch {
		case (c == 'w' || k == 65517) && g.direction != down:
			g.direction = up
		case (c == 's' || k == 65516) && g.direction != up:
			g.direction = down
		case (c == 'd' || k == 65514) && g.direction != left:
			g.direction = right
		case (c == 'a' || k == 65515) && g.direction != right:
			g.direction = left
		case c == 'q':
			return
		}
	}
}

func (g *game) checkStep() {
	switch f := g.snake[0]; {
	case slices.ContainsFunc(g.snake[1:], func(c *cell) bool { return f == c }): // check if snake steps on itself
		panic("Game over")
	case len(g.snake) > (int(math.Pow(float64(g.gridSize), 2) - 2)): // check if there is no free cells for game
		g.gameOver = true
	case f == g.item: // check if snake "eat" bonus item
		g.snake = append(g.snake, g.snakeBackLastPosition)
		g.item = g.getRandomCell()
	}
}

func (g *game) getRandomCell() (c *cell) {
	for {
		x, y := getRandomPoint(g.gridSize - 1)
		c = g.grid[y][x]

		if slices.Contains(g.snake, c) {
			continue
		}

		return
	}
}
