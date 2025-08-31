// Package main is the startup for the supply run api service.
package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/b-sea/go-auth/password"
	"github.com/b-sea/go-auth/password/encrypt"
	"github.com/b-sea/go-auth/token"
	"github.com/b-sea/supply-run-api/internal/auth"
	"github.com/b-sea/supply-run-api/internal/entity"
	"github.com/b-sea/supply-run-api/internal/memory"
	"github.com/b-sea/supply-run-api/internal/mock"
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

	recorder := &mock.Recorder{}

	svr := server.NewServer(
		auth.NewService(
			memory.NewRepository(
				memory.WithAccounts(
					auth.NewAccount(entity.NewID(), "bcarl", "$argon2id$v=19$m=12,t=1,p=3$OlELETZJmMlruz2YivfZvw$sj4d2TYVRYwdB7FH9iYDEvb79hGwAlAqxESfxWnSSNw", time.Now().UTC()),
					auth.NewAccount(entity.NewID(), "someone", "$argon2id$v=19$m=12,t=1,p=3$U8WvHCC/mU8zy7hcEyUy1Q$8gIKD24Bn9RakcSdBByrpmCgR2AIBHnERJXfmOiD8jg", time.Now().UTC()),
				),
			),
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
