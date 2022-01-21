package commands

import (
	"andy/pkg/http_server"
)

// var availableArgs []string = []string{"start", "status", "stop"}

// var Server ICommand = Command{true, availableArgs}

func RunDockerStatusServer() {

	// if len(args) > 1 && args[1] == "-p" {
	// 	if _, err := strconv.Atoi(args[2]); len(args) < 3 || err != nil {
	// 		fmt.Println("⛔️Please, set correct port!")
	// 		os.Exit(0)
	// 	}

	// 	port = args[2]
	// }

	http_server.RunDockerServer()
}
