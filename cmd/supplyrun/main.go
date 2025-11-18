// Package main is the startup for the supply run api service.
package main

import (
	"os"

	"github.com/b-sea/supply-run-api/cmd/supplyrun/cli"
)

var version = "unversioned"

func main() {
	if err := cli.New(version).Execute(); err != nil {
		os.Exit(1)
	}
}
