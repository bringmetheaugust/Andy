package commands

import (
	"andy/pkg/colorFmt"
	"os"

	"github.com/spf13/cobra"
)

const shellFlag = "shell"

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringP(shellFlag, "s", "", "Set completion shell type.")
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init package for more userfriendly usage.",
	Run: func(cmd *cobra.Command, args []string) {
		shellType, _ := cmd.Flags().GetString(shellFlag)

		generateCompletion(cmd, shellType)
	},
}

func generateCompletion(cmd *cobra.Command, shellType string) {
	if shellType == "" {
		colorFmt.YellowFmt("You did't set shell type. I'll generate completion for bash shell for U.")

		shellType = "bash"
	}

	switch shellType {
	case "bash":
		cmd.Root().GenBashCompletion(os.Stdout)
	case "zsh":
		cmd.Root().GenZshCompletion(os.Stdout)
	case "fish":
		cmd.Root().GenFishCompletion(os.Stdout, true)
	case "powershell":
		cmd.Root().GenPowerShellCompletionWithDesc(os.Stdout)
	default:
		colorFmt.RedFmt("U set unknown shell type. I'll skip completion for Your shell.")
	}
}
