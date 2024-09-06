package commands

import (
	"andy/internal/services/graph"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(graphStartCmd)
}

var graphStartCmd = &cobra.Command{
	Use:     "graph",
	Short:   "Build graph📊. Paste a few arguments as numbers between 0 and 10 for each collumn.",
	Example: "graph 8 1 2",
	Run: func(cmd *cobra.Command, args []string) {
		graph.Run(args)
	},
}
