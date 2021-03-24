package commands

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strconv"
)

var availableArgs []string = []string{"start", "status", "stop"}

var Server ICommand = Command{true, availableArgs}

func (c Command) Run(args []string) {
	switch args[0] {
	case "start":
		{
			port := "1991" // default port

			if len(args) > 1 && args[1] == "-p" {
				if _, err := strconv.Atoi(args[2]); len(args) < 3 || err != nil {
					fmt.Println("⛔️Please, set correct port!")
					os.Exit(0)
				}

				port = args[2]
			}

			runServer(port)
		}
	}
}

func runServer(port string) {
	http.HandleFunc("/docker/ps", getDockerStatus)
	http.ListenAndServe(":"+port, nil)

	fmt.Println("Server daemon is running on " + port + " port!")
}

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
