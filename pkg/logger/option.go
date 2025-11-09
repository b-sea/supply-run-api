package logger

import (
	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Option is a creation option for the global zerolog logger.
type Option func(details *setup)

// WithLevel sets the global zerolog logger to a log level.
func WithLevel(level string) Option {
	return func(details *setup) {
		lvl, err := zerolog.ParseLevel(level)
		if err != nil {
			lvl = zerolog.InfoLevel
		}

		details.log = details.log.Level(lvl)
	}
}

// WithRotation sets up log rotation for the global zerolog logger.
func WithRotation(filename string, maxSize int, maxBackups int, maxAge int, compress bool) Option {
	return func(details *setup) {
		details.writers = append(details.writers,
			&lumberjack.Logger{
				Filename:   filename,
				MaxSize:    maxSize,
				MaxBackups: maxBackups,
				MaxAge:     maxAge,
				Compress:   compress,
			},
		)
	}
}
