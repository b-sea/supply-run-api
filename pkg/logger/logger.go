// Package logger manages a global zerolog logger.
package logger

import (
	"io"
	"os"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

var (
	once sync.Once      //nolint: gochecknoglobals
	log  zerolog.Logger //nolint: gochecknoglobals
)

type setup struct {
	log     zerolog.Logger
	writers []io.Writer
}

// Get the global zerolog logger.
func Get() zerolog.Logger {
	Setup()

	return log
}

// Setup the global zerolog logger.
func Setup(options ...Option) {
	once.Do(func() {
		zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack //nolint: reassign
		zerolog.TimeFieldFormat = time.RFC3339Nano

		details := &setup{
			log: zerolog.New(nil).Level(zerolog.InfoLevel),
			writers: []io.Writer{
				zerolog.ConsoleWriter{
					Out:        os.Stdout,
					TimeFormat: time.RFC3339,
				},
			},
		}

		for _, option := range options {
			option(details)
		}

		log = details.log.
			Output(zerolog.MultiLevelWriter(details.writers...)).
			With().Timestamp().
			Logger()
	})
}
