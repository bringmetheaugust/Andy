package commands

import (
	"andy/constants"

	"github.com/spf13/cobra"
)

var CONFIG_FILE string

func init() {
	// cobra.OnInitialize()

	// rootCMD.PersistentFlags().StringVar(&CONFIG_FILE, "config", "", "config file (default is $HOME/.dadjoke.yaml)")
	// rootCMD.Flags().BoolP("toggle", "t", false, "Help message for toggle")
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
