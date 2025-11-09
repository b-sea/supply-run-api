// Package main is the startup for the supply run api service.
package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/b-sea/supply-run-api/internal/metrics"
	"github.com/b-sea/supply-run-api/internal/server"
	"github.com/sirupsen/logrus"
)

func main() {
	svr := server.New(metrics.NewBasicLogger())

	channel := make(chan os.Signal, 1)
	signal.Notify(channel, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := svr.Start(); err != nil {
			logrus.Fatalf("error starting server: %v", err)
		}
	}()

	<-channel

	if err := svr.Stop(); err != nil {
		logrus.Fatalf("server forced to shutdown: %v", err)
	}

	logrus.Info("server gracefully stopped")
}
