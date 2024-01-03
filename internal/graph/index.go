package graph

import "fmt"

const maxRange int = 10

func BuildGraph(nums []int) {
	for _, v := range nums {
		if v > maxRange {
			fmt.Printf("%v is too big.\n", v)

			return
		}
	}

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
