package commands

import (
	"andy/constants"
	"fmt"

	"github.com/spf13/cobra"
)

var CONFIG_FILE string

func init() {
	// cobra.OnInitialize()

	// rootCMD.PersistentFlags().StringVar(&CONFIG_FILE, "config", "", "config file (default is $HOME/.dadjoke.yaml)")
	// rootCMD.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

var rootCMD = &cobra.Command{
	Use:   "andy",
	Short: "As U know, I'm Andy🐼 and I have a lot of funny tools.",
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
	Version: "0.0.1",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hi! I'm Andy🐼!\n" + constants.TooltipHelpMessage)
	},
}

func ParseCMD() {
	rootCMD.Execute()
}
