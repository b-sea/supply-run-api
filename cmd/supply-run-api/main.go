// Package main is the startup for the supply run api service.
package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/b-sea/go-auth/password"
	"github.com/b-sea/go-auth/password/encrypt"
	"github.com/b-sea/go-auth/token"
	"github.com/b-sea/supply-run-api/internal/auth"
	"github.com/b-sea/supply-run-api/internal/memory"
	"github.com/b-sea/supply-run-api/internal/prometheus"
	"github.com/b-sea/supply-run-api/internal/server"
	"github.com/sirupsen/logrus"
)

func main() {
	publicKey, err := os.ReadFile("../../.certs/key.pub")
	if err != nil {
		logrus.Fatal(err)
	}

	privateKey, err := os.ReadFile("../../.certs/key.pem")
	if err != nil {
		logrus.Fatal(err)
	}

	tokenService, err := token.NewService(publicKey, privateKey, token.WithIssuer("supply-run-api"))
	if err != nil {
		logrus.Fatal(err)
	}

	recorder := prometheus.NewRecorder("supply-run-api")

	svr := server.NewServer(
		auth.NewService(
			memory.NewRepository(),
			tokenService,
			password.NewService(encrypt.NewArgon2Repo()),
			recorder,
		),
		recorder,
	)
	svr.Start()

	channel := make(chan os.Signal, 1)
	signal.Notify(channel, syscall.SIGINT, syscall.SIGTERM)

	<-channel

	svr.Stop()
}
