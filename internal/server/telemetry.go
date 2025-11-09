package server

import (
	"net/http"
	"time"

	"github.com/b-sea/supply-run-api/internal/server/graphql"
	"github.com/b-sea/supply-run-api/pkg/logger"
	"github.com/gorilla/mux"
	"github.com/lithammer/shortuuid"
	"github.com/rs/zerolog"
)

// Recorder defines functions for tracking HTTP-based metrics.
type Recorder interface {
	Handler() http.Handler

	ObserveRequestDuration(method string, path string, code int, duration time.Duration)
	ObserveResponseSize(method string, path string, code int, bytes int64)

	graphql.Recorder
}

type telemetryWriter struct {
	http.ResponseWriter

	StatusCode int
	Size       int
}

func (w *telemetryWriter) WriteHeader(statusCode int) {
	w.StatusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func (w *telemetryWriter) Write(p []byte) (int, error) {
	w.Size += len(p)

	return w.ResponseWriter.Write(p) //nolint: wrapcheck
}

func telemetryMiddleware(recorder Recorder) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			start := time.Now()
			log := logger.Get()

			path, err := mux.CurrentRoute(request).GetPathTemplate()
			if err != nil {
				path = request.URL.Path
			}

			hijack := &telemetryWriter{
				ResponseWriter: writer,
				StatusCode:     http.StatusOK,
				Size:           0,
			}

			defer func() {
				duration := time.Since(start)

				log.Info().
					Str("method", request.Method).
					Str("url", request.URL.RequestURI()).
					Str("user_agent", request.UserAgent()).
					Int("status_code", hijack.StatusCode).
					Dur("duration_ms", duration).
					Int("response_bytes", hijack.Size).
					Msg("http request")

				recorder.ObserveRequestDuration(request.Method, path, hijack.StatusCode, duration)
				recorder.ObserveResponseSize(request.Method, path, hijack.StatusCode, int64(hijack.Size))
			}()

			// Add a correlation ID
			correlationID := shortuuid.New()
			hijack.Header().Add("Correlation-ID", correlationID)
			log.UpdateContext(func(c zerolog.Context) zerolog.Context {
				return c.Str("correlation_id", correlationID)
			})

			next.ServeHTTP(hijack, request.WithContext(logger.Get().WithContext(request.Context())))
		})
	}
}
