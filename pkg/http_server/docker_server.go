package http_server

import (
	"fmt"
	"net/http"
	"os/exec"
	"strconv"
)

func RunDockerHttpServer(port int64) {
	if port == 0 {
		port = 1991
	}

	var httpPort string = strconv.FormatInt(port, 10)

	fmt.Println("Trying to star HTTP docker server on " + httpPort + " port ...")

	http.HandleFunc("/ps", getDockerStatus)
	http.ListenAndServe(":"+httpPort, nil)

	fmt.Println("Server daemon is running on " + httpPort + " port!")
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
