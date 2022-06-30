package main

import (
	"fmt"
	"os"

	"github.com/stackitcloud/stackit-api-manager-cli/cmd/stackit/cmd"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "an error occurred: %s\n", err)
		os.Exit(1)
	}
}

func run() error {
	cmd.Execute()

	return nil
}
