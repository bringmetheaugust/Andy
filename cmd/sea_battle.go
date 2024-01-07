package commands

import (
	seabattle "andy/internal/services/sea_battle"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(seaBattleCMD)
	seaBattleCMD.AddCommand(seaBattleStartCmd)
}

var seaBattleCMD = &cobra.Command{
	Use:   "seabattle",
	Short: "Play sea battle game.",
}

var seaBattleStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start game.",
	Run: func(cmd *cobra.Command, args []string) {
		seabattle.StartGame()
	},
}
