package commands

import (
	"github.com/spf13/cobra"
)

var CONFIG_FILE string

func init() {
	// cobra.OnInitialize()
}

var rootCmd = &cobra.Command{
	Use:   "andy",
	Short: "Hi! I'm Andy🐼!",
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
	Version: "0.0.1",
}

func ParseCmd() {
	rootCmd.Execute()
}
