package cli

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/b-sea/go-config/config"
	"github.com/b-sea/go-server/server"
	"github.com/b-sea/supply-run-api/internal/graphql"
	"github.com/b-sea/supply-run-api/internal/mariadb"
	"github.com/b-sea/supply-run-api/internal/metrics"
	"github.com/b-sea/supply-run-api/internal/query"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"github.com/spf13/cobra"
)

func startCmd(version string) *cobra.Command {
	cmd := &cobra.Command{
		Version: version,
		Use:     "start config",
		Short:   "start the Supply Run API web service",
		Args:    cobra.ExactArgs(1),
		RunE:    startRun(),
	}

	return cmd
}

func startRun() func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		cfg := defaultConfig()

		if err := config.Load(&cfg, config.WithEnvPrefix(cfgEnvPrefix), config.WithFile(args[0])); err != nil {
			return err
		}

		log := setupLogger(cfg)
		recorder := metrics.NewPrometheus()
		repository := mariadb.NewRepository(
			mariadb.BasicConnector(cfg.MariaDB.Host, cfg.MariaDB.Username, cfg.MariaDB.Password),
		)

		svr := server.New(log, recorder,
			server.SetPort(cfg.Server.Port),
			server.SetReadTimeout(time.Duration(cfg.Server.ReadTimeout)*time.Second),
			server.SetWriteTimeout(time.Duration(cfg.Server.WriteTimeout)*time.Second),
			server.SetVersion(cmd.Version),
			server.AddHandler(
				"/graphql",
				graphql.New(
					query.NewService(repository, repository, repository),
					recorder,
				),
				http.MethodPost,
			),
			server.AddHealthDependency("database", repository),
		)

		channel := make(chan os.Signal, 1)
		signal.Notify(channel, syscall.SIGINT, syscall.SIGTERM)

		go func() {
			if err := svr.Start(); err != nil {
				log.Error().Err(err).Msg("error starting server")
			}
		}()

		<-channel

		if err := svr.Stop(); err != nil {
			log.Error().Err(err).Msg("server forced to shutdown")

			return err
		}

		log.Info().Msgf("server gracefully stopped")

		return nil
	}
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
