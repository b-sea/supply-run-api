// Package main is the startup for the supply run api service.
package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/b-sea/supply-run-api/internal/metrics"
	"github.com/b-sea/supply-run-api/internal/server"
	"github.com/b-sea/supply-run-api/pkg/logger"
)

func main() {
	log := logger.Get()
	svr := server.New(metrics.NewNoOp())

	channel := make(chan os.Signal, 1)
	signal.Notify(channel, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := svr.Start(); err != nil {
			log.Fatal().Err(err).Msg("error starting server")
		}
	}()

	<-channel

	if err := svr.Stop(); err != nil {
		log.Fatal().Err(err).Msg("server forced to shutdown")
	}

	log.Info().Msgf("server gracefully stopped")
}
