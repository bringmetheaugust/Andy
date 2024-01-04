package commands

import (
	"andy/internal/graph"
	"andy/pkg/colorFmt"
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
			colorFmt.YellowFmt("There are no arguments. Please, pass some agruments.")

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
