package commands

import (
	"andy/internal/services/docker_server"

	"github.com/spf13/cobra"
)

const portFlag = "port"

func init() {
	RootCmd.AddCommand(dockerHTTPCmd)
	dockerHTTPCmd.AddCommand(dockerHTTPStartCmd)
	dockerHTTPStartCmd.Flags().Int64P(portFlag, "p", 1991, "Set HTTP server port")
}

var dockerHTTPCmd = &cobra.Command{
	Use:   "docker_http",
	Short: "Run HTTP server to get local docker daemon information🐳.",
}

var dockerHTTPStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Run server",
	Run: func(cmd *cobra.Command, args []string) {
		port, _ := cmd.Flags().GetInt64(portFlag)

		docker_server.RunDockerHttpServer(port)
	},
}
