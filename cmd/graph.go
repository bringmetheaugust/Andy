package commands

import (
	"andy/internal/services/graph"
	"andy/pkg/colorPrint"
	"strconv"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(graphStartCmd)
}

var graphStartCmd = &cobra.Command{
	Use:     "graph",
	Short:   "Build graph. Paste a few arguments as numbers between 0 and 10 for each collumn.",
	Example: "graph 8 1 2",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			colorPrint.RedPrint("There are no arguments. Please, pass some agruments.\n")

			return
		}

		var mapedArr []uint8

		for _, v := range args {
			var fInt uint8
			var i, _ = strconv.Atoi(v)

			fInt = uint8(i)
			mapedArr = append(mapedArr, fInt)
		}

		graph.BuildGraph(mapedArr)
	},
}
