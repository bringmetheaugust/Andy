package docker_server

import (
	"andy/pkg/colorFmt"
	"net/http"
	"os/exec"
	"strconv"
)

func RunDockerHttpServer(port int64) {
	if port == 0 {
		port = 1991
	}

	var httpPort string = strconv.FormatInt(port, 10)

	colorFmt.BlueFmt("Trying to start HTTP docker server on " + httpPort + " port...")

	http.HandleFunc("/ps", getDockerStatus)
	http.ListenAndServe(":"+httpPort, nil)

	colorFmt.GreenFmt("Server daemon has been successfully started!")
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
