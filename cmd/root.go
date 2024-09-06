package commands

import (
	"andy/constants"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Short: "Hi! I'm Andy" + constants.Logo + "!",
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
	Version: "0.0.3",
}
