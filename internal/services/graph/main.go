package graph

import (
	"andy/pkg/colorPrint"
	"errors"
	"fmt"
	"strconv"
)

const maxRange uint8 = 10

type args []string
type columns []uint8

func Run(a args) {
	c, err := requestGraphData(a)

	if err != nil {
		colorPrint.RedPrint(err.Error())

		return
	}

	buildGraph(c)
}

func requestGraphData(a args) (columns, error) {
	if len(a) == 0 {
		return nil, errors.New("there are no arguments. Please, pass some agruments")
	}

	c := make(columns, len(a))

	for _, v := range a {
		var i, err = strconv.Atoi(v)

		if err != nil {
			return nil, fmt.Errorf("invalid %v argument. Please, type correct number", v)
		}

		c = append(c, uint8(i))
	}

	for _, v := range c {
		if v > maxRange {
			return nil, fmt.Errorf("%d is too big. Please, paste arguments betwwen 0 and %d", v, maxRange)
		}
	}

	return c, nil
}

func buildGraph(nums columns) {
	for r := maxRange; r > 0; r-- {
		for _, v := range nums {
			switch {
			case r > v:
				fmt.Printf("***")
			case r < v:
				colorPrint.BluePrint("| |")
			default:
				colorPrint.BluePrint("╭─╮")
			}
		}

		fmt.Printf(" %v\n", r)
	}
}
