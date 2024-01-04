package main

import (
	commands "andy/cmd"
	"fmt"
	"os"
)

// func init() {
// cobra.OnInitialize()
// }

func main() {
	if err := commands.RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
