// Package server implements the API web server.
package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/b-sea/supply-run-api/internal/graphql"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
)

const (
	defaultPort    = 5000
	defaultTimeout = 10 * time.Second
)

// Server is a supply-run API web server.
type Server struct {
	http      *http.Server
	validator *validator.Validate
	log       zerolog.Logger
}

// New creates a new Server.
func New(log zerolog.Logger, recorder Recorder, options ...Option) *Server {
	server := &Server{
		http: &http.Server{
			Addr:              fmt.Sprintf(":%d", defaultPort),
			ReadTimeout:       defaultTimeout,
			ReadHeaderTimeout: defaultTimeout,
			WriteTimeout:      defaultTimeout,
		},
		validator: validator.New(),
		log:       log,
	}

	for _, option := range options {
		option(server)
	}

	router := mux.NewRouter()
	router.Use(telemetryMiddleware(log, recorder))

	router.Handle(
		"/metrics",
		func() http.HandlerFunc {
			return func(writer http.ResponseWriter, request *http.Request) {
				recorder.Handler().ServeHTTP(writer, request)
			}
		}(),
	).Methods(http.MethodGet)

	router.Handle("/graphql", graphql.NewHandler(recorder)).Methods(http.MethodPost)

	// Re-define the default NotFound handler so it passes through middleware correctly.
	router.NotFoundHandler = router.NewRoute().HandlerFunc(http.NotFound).GetHandler()
	server.http.Handler = router

	_ = router.Walk(func(route *mux.Route, _ *mux.Router, _ []*mux.Route) error {
		if route.GetHandler() == nil {
			return nil
		}

		template, err := route.GetPathTemplate()
		if err != nil {
			return nil //nolint: nilerr
		}

		methods, _ := route.GetMethods()
		if len(methods) == 0 {
			log.Debug().Str("url", template).Msg("route registered")

			return nil
		}

		slices.Sort(methods)

		for i := range methods {
			log.Debug().Str("method", methods[i]).Str("url", template).Msg("route registered")
		}

		return nil
	})

	return server
}

// Addr returns the server address.
func (s *Server) Addr() string {
	return s.http.Addr
}

// ReadTimeout returns the server read and read header timeout.
func (s *Server) ReadTimeout() time.Duration {
	return s.http.ReadTimeout
}

// WriteTimeout returns the server write timeout.
func (s *Server) WriteTimeout() time.Duration {
	return s.http.WriteTimeout
}

func (s *Server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	s.http.Handler.ServeHTTP(writer, request)
}

// Start the Server.
func (s *Server) Start() error {
	s.log.Info().Str("addr", s.http.Addr).Msg("starting server")

	if err := s.http.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		return err //nolint: wrapcheck
	}

	return nil
}

// Stop the Server.
func (s *Server) Stop() error {
	s.log.Info().Str("addr", s.http.Addr).Msg("stopping server")

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	return s.http.Shutdown(ctx) //nolint: wrapcheck
}
