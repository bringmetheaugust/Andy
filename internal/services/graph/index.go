package graph

import (
	"andy/internal/components"
	"andy/pkg/colorPrint"
	"fmt"
	"strconv"
	"time"
)

const maxRange uint8 = 10

// Print graph to CLI
func BuildGraph(nums []uint8) {
	for _, v := range nums {
		if v > maxRange || v < 0 {
			colorPrint.RedPrint(strconv.Itoa(int(v)) + " is too big. Please, paste arguments betwwen 0 and " + strconv.Itoa(int(maxRange)) + ".\n")

			return
		}
	}

	components.Spinner.Start()
	time.Sleep(1 * time.Second)
	components.Spinner.Stop()

	for raw := maxRange; raw > 0; raw-- {
		for _, v := range nums {
			switch {
			case raw > v:
				fmt.Printf("***")
			case raw < v:
				colorPrint.BluePrint("| |")
			default:
				colorPrint.BluePrint("╭─╮")
			}
		}

		fmt.Printf(" %v\n", raw)
	}
}
