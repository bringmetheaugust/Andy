package http_server

import (
	"fmt"
	"net/http"
	"os/exec"
)

// http server to manage docker containers
func RunDockerServer() {
	port := "1991" // default port

	http.HandleFunc("/ps", getDockerStatus)
	http.ListenAndServe(":"+port, nil)

	fmt.Println("Server daemon is running on " + port + " port!")
}

// get docker container's statuses
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
