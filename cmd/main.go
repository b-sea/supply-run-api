// Package main is the startup for the supply run api service.
package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/b-sea/supply-run-api/internal/metrics"
	"github.com/b-sea/supply-run-api/internal/server"
)

func main() {
	svr := server.NewServer(metrics.NewBasicLogger())
	svr.Start()

	channel := make(chan os.Signal, 1)
	signal.Notify(channel, syscall.SIGINT, syscall.SIGTERM)

	<-channel

	svr.Stop()
}
