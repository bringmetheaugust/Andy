package commands

import (
	seabattle "andy/internal/services/sea_battle"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(seaBattleCMD)
	seaBattleCMD.AddCommand(seaBattleStartCmd)
	seaBattleCMD.AddCommand(seaBattleStartDevCmd)
}

var seaBattleCMD = &cobra.Command{
	Use:   "seabattle",
	Short: "Play sea battle game⚓.",
}

var seaBattleStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start game.",
	Run: func(cmd *cobra.Command, args []string) {
		seabattle.StartGame(false)
	},
}

var seaBattleStartDevCmd = &cobra.Command{
	Use:   "dev",
	Short: "Start game with development mode.",
	Run: func(cmd *cobra.Command, args []string) {
		seabattle.StartGame(true)
	},
}
