package graph

import (
	"andy/internal/components"
	"fmt"
	"time"
)

const maxRange int = 10

func BuildGraph(nums []int) {
	for _, v := range nums {
		if v > maxRange {
			fmt.Printf("%v is too big.\n", v)

			return
		}
	}

	components.Spinner.Start()
	time.Sleep(2 * time.Second)
	components.Spinner.Stop()

	for count := maxRange; count > 0; count-- {
		for _, v := range nums {
			switch {
			case count > v:
				fmt.Printf("***")
			case count < v:
				fmt.Printf("| |")
			default:
				fmt.Printf("---")
			}
		}

		fmt.Printf(" %v\n", count)
	}
}
