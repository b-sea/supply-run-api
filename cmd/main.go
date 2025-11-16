// Package main is the startup for the supply run api service.
package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/b-sea/go-config/config"
	"github.com/b-sea/go-server/server"
	"github.com/b-sea/supply-run-api/internal/graphql"
	"github.com/b-sea/supply-run-api/internal/metrics"
	"github.com/b-sea/supply-run-api/internal/mock"
	"github.com/b-sea/supply-run-api/internal/query"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

const cfgEnvPrefix = "SUPPLYRUN"

var version string

func main() {
	options := []config.Option{
		config.WithEnvPrefix(cfgEnvPrefix),
	}

	cfgFile := configFile()
	if cfgFile != "" {
		options = append(options, config.WithFile(cfgFile))
	}

	cfg := defaultConfig()

	if err := config.Load(&cfg, options...); err != nil {
		panic(err)
	}

	log := setupLogger(cfg)
	recorder := metrics.NewNoOp()

	svr := server.New(log, recorder,
		server.SetPort(cfg.Server.Port),
		server.SetReadTimeout(time.Duration(cfg.Server.ReadTimeout)*time.Second),
		server.SetWriteTimeout(time.Duration(cfg.Server.WriteTimeout)*time.Second),
		server.SetVersion(version),
		server.AddHandler("/graphql", graphql.New(query.NewService(&mock.QueryRepository{}), recorder), http.MethodPost),
	)

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

func setupLogger(cfg Config) zerolog.Logger {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack //nolint: reassign
	zerolog.TimeFieldFormat = time.RFC3339Nano

	level, err := zerolog.ParseLevel(cfg.Logger.Level)
	if err != nil {
		level = zerolog.InfoLevel
	}

	writer := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	}

	log := zerolog.New(writer).Level(level).With().Timestamp().Logger()
	zerolog.DefaultContextLogger = &log

	return log
}

func configFile() string {
	configFile, _ := os.LookupEnv(cfgEnvPrefix + "_CONFIGFILE")

	return configFile
}

type Config struct {
	Server struct {
		Port         int `config:"port"`
		ReadTimeout  int `config:"readTimeout"`
		WriteTimeout int `config:"writeTimeout"`
	} `config:"server"`

	Logger struct {
		Level string `config:"level"`
	} `config:"logger"`
}

func defaultConfig() Config {
	return Config{
		Server: struct {
			Port         int `config:"port"`
			ReadTimeout  int `config:"readTimeout"`
			WriteTimeout int `config:"writeTimeout"`
		}{
			Port:         5000,
			ReadTimeout:  5,
			WriteTimeout: 5,
		},
		Logger: struct {
			Level string `config:"level"`
		}{
			Level: "debug",
		},
	}
}
