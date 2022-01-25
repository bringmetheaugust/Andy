package commands

import (
	"andy/constants"

	"github.com/spf13/cobra"
)

var CONFIG_FILE string

func init() {
	// cobra.OnInitialize()
}

var rootCmd = &cobra.Command{
	Use:   "andy",
	Short: "Hi! I'm Andy🐼!\n" + constants.TooltipHelpMessage,
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
	Version: "0.0.1",
}

func ParseCmd() {
	rootCmd.Execute()
}
