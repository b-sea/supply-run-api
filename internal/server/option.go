package server

import (
	"fmt"
	"net/http"
	"slices"
	"time"
)

// Option is a creation option for a Server.
type Option func(server *Server)

// WithPort overrides the port used by the Server.
func WithPort(port int) Option {
	return func(server *Server) {
		if port == defaultPort {
			return
		}

		server.log.Debug().Int("port", port).Msgf("override server port")
		server.http.Addr = fmt.Sprintf(":%d", port)
	}
}

// WithReadTimeout overrides the HTTP read and read header timeouts for the Server.
func WithReadTimeout(duration time.Duration) Option {
	return func(server *Server) {
		if duration <= 0 || duration == defaultTimeout {
			return
		}

		server.log.Debug().Dur("timeout_ms", duration).Msg("override server read timeout")
		server.http.ReadTimeout = duration
		server.http.ReadHeaderTimeout = duration
	}
}

// WithWriteTimeout overrides the HTTP write timeout for the Server.
func WithWriteTimeout(duration time.Duration) Option {
	return func(server *Server) {
		if duration <= 0 || duration == defaultTimeout {
			return
		}

		server.log.Debug().Dur("timeout_ms", duration).Msg("override server write timeout")
		server.http.WriteTimeout = duration
	}
}

// AddHandler adds an HTTP hander to the Server.
func AddHandler(path string, handler http.Handler, methods ...string) Option {
	return func(server *Server) {
		if len(methods) == 0 {
			server.log.Debug().Str("path", path).Msg("route registered")
		}

		slices.Sort(methods)

		for i := range methods {
			server.log.Debug().Str("method", methods[i]).Str("path", path).Msg("route registered")
		}

		server.router.Handle(path, handler).Methods(methods...)
	}
}
