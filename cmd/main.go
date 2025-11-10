// Package main is the startup for the supply run api service.
package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/b-sea/go-config/config"
	"github.com/b-sea/go-logger/logger"
	"github.com/b-sea/supply-run-api/internal/metrics"
	"github.com/b-sea/supply-run-api/internal/server"
)

const cfgEnvPrefix = "SUPPLYRUN"

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

	logger.Setup(
		logger.WithLevel(cfg.Logger.Level),
		logger.WithRotation(
			cfg.Logger.Rotation.FilePath,
			cfg.Logger.Rotation.MaxSize,
			cfg.Logger.Rotation.MaxAge,
			cfg.Logger.Rotation.MaxBackups,
			cfg.Logger.Rotation.Compress,
		),
	)

	log := logger.Get()

	svr := server.New(
		metrics.NewNoOp(),
		server.WithPort(cfg.Server.Port),
		server.WithReadTimeout(time.Duration(cfg.Server.ReadTimeout)*time.Second),
		server.WithWriteTimeout(time.Duration(cfg.Server.WriteTimeout)*time.Second),
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
		Level    string `config:"level"`
		Rotation struct {
			FilePath   string `config:"filePath"`
			MaxSize    int    `config:"maxSize"`
			MaxBackups int    `config:"maxBackups"`
			MaxAge     int    `config:"maxAge"`
			Compress   bool   `config:"compress"`
		} `config:"rotation"`
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
			Level    string `config:"level"`
			Rotation struct {
				FilePath   string `config:"filePath"`
				MaxSize    int    `config:"maxSize"`
				MaxBackups int    `config:"maxBackups"`
				MaxAge     int    `config:"maxAge"`
				Compress   bool   `config:"compress"`
			} `config:"rotation"`
		}{
			Level: "debug",
			Rotation: struct {
				FilePath   string `config:"filePath"`
				MaxSize    int    `config:"maxSize"`
				MaxBackups int    `config:"maxBackups"`
				MaxAge     int    `config:"maxAge"`
				Compress   bool   `config:"compress"`
			}{
				FilePath:   "./log/supply-run-api.log",
				MaxSize:    10,
				MaxBackups: 5,
				MaxAge:     14,
				Compress:   true,
			},
		},
	}
}
