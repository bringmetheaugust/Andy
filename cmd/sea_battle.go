package commands

import (
	seabattle "andy/internal/services/sea_battle"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(seaBattleStartCmd)
}

var seaBattleStartCmd = &cobra.Command{
	Use:   "seabattle",
	Short: "Play sea battle game",
	Run: func(cmd *cobra.Command, args []string) {
		seabattle.StartGame()
	},
}
