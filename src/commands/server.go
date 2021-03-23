package commands

import (
	"andy/src/utils"
	"net/http"
	"os/exec"
)

var availableArgs []string = []string{"start", "status", "stop"}
var Server Command = Command{true, availableArgs}

func getDockerStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")

	cmd := exec.Command("docker", "ps", "-a", "--format", "table {{ .ID }}\t{{ .Names }}\t{{.Status}}\t{{.Ports}}")
	stdout, err := cmd.Output()

	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)

		return
	}

	w.Write(stdout)
}

func runServer(args []string) {
	port := "1991" // default port

	if len(args) == 0 {
		utils.LessArgument()

		return
	}

	http.HandleFunc("/docker/ps", getDockerStatus)
	http.ListenAndServe(":"+port, nil)
}
