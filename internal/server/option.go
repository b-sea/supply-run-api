package server

import (
	"fmt"
	"time"

	"github.com/b-sea/go-logger/logger"
)

// Option is a creation option for a Server.
type Option func(server *Server)

// WithPort overrides the port used by the Server.
func WithPort(port int) Option {
	return func(server *Server) {
		log := logger.Get()

		if port == 0 {
			log.Warn().Msgf("server port cannot be set to %d, defaulting to %d", port, defaultPort)
		}

		if port == defaultPort {
			return
		}

		log.Debug().Int("port", port).Msgf("override server port")
		server.http.Addr = fmt.Sprintf(":%d", port)
	}
}

// WithReadTimeout overrides the HTTP read and read header timeouts for the Server.
func WithReadTimeout(duration time.Duration) Option {
	return func(server *Server) {
		log := logger.Get()

		if duration == 0 || duration == defaultTimeout {
			return
		}

		log.Debug().Dur("timeout_ms", duration).Msg("override server read timeout")
		server.http.ReadTimeout = duration
		server.http.ReadHeaderTimeout = duration
	}
}

// WithWriteTimeout overrides the HTTP write timeout for the Server.
func WithWriteTimeout(duration time.Duration) Option {
	return func(server *Server) {
		log := logger.Get()

		if duration == 0 || duration == defaultTimeout {
			return
		}

		log.Debug().Dur("timeout_ms", duration).Msg("override server write timeout")
		server.http.WriteTimeout = duration
	}
}
