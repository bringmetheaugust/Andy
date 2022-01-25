package commands

import (
	"andy/pkg/http_server"

	"github.com/spf13/cobra"
)

func init() {
	rootCMD.AddCommand(dockerServerCmd)
}

var dockerServerCmd = &cobra.Command{
	Use:   "docker",
	Short: "Run HTTP server to get local docker daemon information.",
	// Long:  `This command fetches a random dad joke from the icanhazdadjoke api`,
	Run: func(cmd *cobra.Command, args []string) {
		http_server.RunDockerServer()
	},
}
