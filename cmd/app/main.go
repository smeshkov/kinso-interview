package main

import (
	"fmt"
	"os"

	"github.com/smeshkov/kinso-interview/app"
	"github.com/smeshkov/kinso-interview/app/config"
)

var (
	version       = "untagged"
	envName       string
	instanceGroup string
)

func main() {
	run := config.NewRuntime(
		version,
		envName,
		instanceGroup,
	)
	if err := app.Run(run); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
