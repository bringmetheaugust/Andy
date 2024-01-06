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
	Use:   "graph",
	Short: "Build graph",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			colorPrint.YellowPrint("There are no arguments. Please, pass some agruments.")

			return
		}

		var mapedArr []int

		for _, v := range args {
			var ll, _ = strconv.Atoi(v)
			mapedArr = append(mapedArr, ll)
		}

		graph.BuildGraph(mapedArr)
	},
}
