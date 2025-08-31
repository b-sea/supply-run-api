package server

import (
	"net/http"
	"strconv"
	"time"

	"github.com/b-sea/supply-run-api/internal/auth"
	"github.com/b-sea/supply-run-api/internal/server/graphql"
	"github.com/gorilla/mux"
)

// Recorder defines functions for tracking HTTP-based metrics.
type Recorder interface {
	Handler() http.Handler

	ObserveRequestDuration(method string, path string, code string, duration time.Duration)
	ObserveResponseSize(method string, path string, code string, bytes int64)

	graphql.Recorder
}

type metricsWriter struct {
	http.ResponseWriter

	StatusCode int
	Size       int
}

func (w *metricsWriter) WriteHeader(statusCode int) {
	w.StatusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func (w *metricsWriter) Write(p []byte) (int, error) {
	w.Size += len(p)

	return w.ResponseWriter.Write(p) //nolint: wrapcheck
}

func (s *Server) metricsMiddleware() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			path, err := mux.CurrentRoute(request).GetPathTemplate()
			if err != nil {
				path = request.URL.Path
			}

			hijack := &metricsWriter{
				ResponseWriter: writer,
				StatusCode:     http.StatusOK,
				Size:           0,
			}

			start := time.Now()
			defer func() {
				code := strconv.Itoa(hijack.StatusCode)

				s.recorder.ObserveRequestDuration(request.Method, path, code, time.Since(start))
				s.recorder.ObserveResponseSize(request.Method, path, code, int64(hijack.Size))
			}()

			next.ServeHTTP(hijack, request)
		})
	}
}

func (s *Server) metricsHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if _, err := auth.FromContext(request.Context()); err != nil {
			auth.Challenge(writer)
			return
		}

		s.recorder.Handler().ServeHTTP(writer, request)
	}
}
