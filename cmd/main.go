// Main package is the entrypoint for the program
package main

import (
	config "github.com/b-sea/supply-run-api/configs"
	"github.com/b-sea/supply-run-api/internal/data/couchbase"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		logrus.Fatal(err)
	}

	_, err = couchbase.NewConnection(*cfg)
	if err != nil {
		logrus.Fatal(err)
	}
}
