// Package server implements the API web server.
package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/b-sea/supply-run-api/internal/auth"
	"github.com/b-sea/supply-run-api/internal/server/graphql"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

const (
	defaultPort    = 5000
	defaultTimeout = time.Minute
)

// Option is a creation option for a Server.
type Option func(server *Server)

// WithPort overrides the port used by the Server.
func WithPort(port int) Option {
	return func(server *Server) {
		server.http.Addr = fmt.Sprintf(":%d", port)
	}
}

// WithTimeout overrides the HTTP read and write timeouts for the Server.
func WithTimeout(duration time.Duration) Option {
	return func(server *Server) {
		server.http.ReadTimeout = duration
		server.http.ReadHeaderTimeout = duration
		server.http.WriteTimeout = duration
	}
}

// Server is a supply-run API web server.
type Server struct {
	http      *http.Server
	router    *mux.Router
	validator *validator.Validate
	auth      *auth.Service
	recorder  Recorder
}

// NewServer creates a new Server.
func NewServer(auth *auth.Service, recorder Recorder, options ...Option) *Server {
	server := &Server{
		http: &http.Server{
			Addr:              fmt.Sprintf(":%d", defaultPort),
			ReadTimeout:       defaultTimeout,
			ReadHeaderTimeout: defaultTimeout,
			WriteTimeout:      defaultTimeout,
		},
		router:    mux.NewRouter(),
		validator: validator.New(),
		recorder:  recorder,
		auth:      auth,
	}

	for _, option := range options {
		option(server)
	}

	server.router.Use(server.metricsMiddleware(), server.auth.Middleware())
	server.router.Handle("/metrics", server.metricsHandler()).Methods(http.MethodGet)
	server.router.Handle("/graphql", graphql.NewHandler(server.auth, recorder)).Methods(http.MethodPost)
	server.router.HandleFunc("/playground", playground.Handler("Supply Run Playground", "/graphql"))

	// Re-define the default NotFound handler so it passes through middleware correctly.
	server.router.NotFoundHandler = server.router.NewRoute().HandlerFunc(http.NotFound).GetHandler()
	server.http.Handler = server.router

	return server
}

// Start the Server.
func (s *Server) Start() {
	go func() {
		logrus.Infof("Starting server %s", s.http.Addr)

		if err := s.http.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}()
}

// Stop the Server.
func (s *Server) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	if err := s.http.Shutdown(ctx); err != nil {
		panic(err)
	}
}
