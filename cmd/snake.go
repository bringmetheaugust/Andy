package commands

import (
	"andy/internal/services/snake"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(snakeCMD)
	snakeCMD.AddCommand(snakeStartCmd)
}

var snakeCMD = &cobra.Command{
	Use:   "snake",
	Short: "Play snake game🐍.",
}

var snakeStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start game.",
	Run: func(cmd *cobra.Command, args []string) {
		snake.StartGame()
	},
}
