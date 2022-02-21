package main

import (
	"fmt"
	"os"

	"wait"
)

func main() {
	url := os.Args[1]
	if err := wait.WaitForServer(url); err != nil {
		fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)
		os.Exit(1)
	}

}
