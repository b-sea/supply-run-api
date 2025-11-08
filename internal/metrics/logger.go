// Package metrics implements all metric recorders.
package metrics

import (
	"net/http"
	"time"

	"github.com/b-sea/supply-run-api/internal/server"
	"github.com/sirupsen/logrus"
)

var (
	_ server.Recorder = (*BasicLogger)(nil)
)

// BasicLogger is a simple metrics recorder that outputs metrics to a log.
type BasicLogger struct{}

// NewBasicLogger creates a new BasicLogger.
func NewBasicLogger() *BasicLogger {
	return &BasicLogger{}
}

// Handler is the logger response to an HTTP request.
func (r *BasicLogger) Handler() http.Handler {
	return http.HandlerFunc(func(_ http.ResponseWriter, _ *http.Request) {})
}

// ObserveRequestDuration records the duration of an HTTP request.
func (r *BasicLogger) ObserveRequestDuration(method string, path string, code string, duration time.Duration) {
	logrus.Infof("Request Duration: %s %s [%s] | %s", method, path, code, duration)
}

// ObserveResponseSize records how large an HTTP response is.
func (r *BasicLogger) ObserveResponseSize(method string, path string, code string, bytes int64) {
	logrus.Infof("Response Size: %s %s [%s] | %v", method, path, code, bytes)
}

// ObserveResolverDuration records the duration of a GraphQL resolver.
func (r *BasicLogger) ObserveResolverDuration(object string, field string, status string, duration time.Duration) {
	logrus.Infof("Resolver Duration: %s.%s [%s] %s", object, field, status, duration)
}

// ObserveResolverError records a GraphQL resolver error.
func (r *BasicLogger) ObserveResolverError(object string, field string, code string) {
	logrus.Infof("Resolver Error: %s.%s | %s", object, field, code)
}
