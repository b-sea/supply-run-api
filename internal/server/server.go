// Package server implements the API web server.
package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/b-sea/supply-run-api/internal/server/graphql"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

const (
	defaultPort    = 5000
	defaultTimeout = 10 * time.Second
)

// Option is a creation option for a Server.
type Option func(server *Server)

// WithPort overrides the port used by the Server.
func WithPort(port int) Option {
	return func(server *Server) {
		server.http.Addr = fmt.Sprintf(":%d", port)
	}
}

// WithReadTimeout overrides the HTTP read and read header timeouts for the Server.
func WithReadTimeout(duration time.Duration) Option {
	return func(server *Server) {
		server.http.ReadTimeout = duration
		server.http.ReadHeaderTimeout = duration
	}
}

// WithWriteTimeout overrides the HTTP write timeout for the Server.
func WithWriteTimeout(duration time.Duration) Option {
	return func(server *Server) {
		server.http.WriteTimeout = duration
	}
}

// Server is a supply-run API web server.
type Server struct {
	http      *http.Server
	validator *validator.Validate
}

// New creates a new Server.
func New(recorder Recorder, options ...Option) *Server {
	server := &Server{
		http: &http.Server{
			Addr:              fmt.Sprintf(":%d", defaultPort),
			ReadTimeout:       defaultTimeout,
			ReadHeaderTimeout: defaultTimeout,
			WriteTimeout:      defaultTimeout,
			Handler:           mux.NewRouter(),
		},
		validator: validator.New(),
	}

	for _, option := range options {
		option(server)
	}

	router := mux.NewRouter()
	router.Use(metricsMiddleware(recorder))

	router.Handle(
		"/metrics",
		func() http.HandlerFunc {
			return func(writer http.ResponseWriter, request *http.Request) {
				recorder.Handler().ServeHTTP(writer, request)
			}
		}(),
	).Methods(http.MethodGet)

	api := router.PathPrefix("/api").Subrouter()
	api.Handle("/graphql", graphql.NewHandler(recorder)).Methods(http.MethodPost)

	// Re-define the default NotFound handler so it passes through middleware correctly.
	router.NotFoundHandler = router.NewRoute().HandlerFunc(http.NotFound).GetHandler()

	server.http.Handler = router

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
	logrus.Infof("Starting server %s", s.http.Addr)

	if err := s.http.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		return err //nolint: wrapcheck
	}

	return nil
}

// Stop the Server.
func (s *Server) Stop() error {
	logrus.Infof("Stopping server %s", s.http.Addr)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	return s.http.Shutdown(ctx) //nolint: wrapcheck
}
