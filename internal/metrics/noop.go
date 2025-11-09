// Package metrics implements all metric recorders.
package metrics

import (
	"net/http"
	"time"

	"github.com/b-sea/supply-run-api/internal/server"
)

var (
	_ server.Recorder = (*NoOp)(nil)
)

// NoOp is a simple metrics recorder that does nothing.
type NoOp struct{}

// NewNoOp creates a new NoOp.
func NewNoOp() *NoOp {
	return &NoOp{}
}

// Handler is the logger response to an HTTP request.
func (r *NoOp) Handler() http.Handler {
	return http.HandlerFunc(func(http.ResponseWriter, *http.Request) {})
}

// ObserveRequestDuration records the duration of an HTTP request.
func (r *NoOp) ObserveRequestDuration(string, string, int, time.Duration) {
}

// ObserveResponseSize records how large an HTTP response is.
func (r *NoOp) ObserveResponseSize(string, string, int, int64) {}

// ObserveResolverDuration records the duration of a GraphQL resolver.
func (r *NoOp) ObserveResolverDuration(string, string, string, time.Duration) {
}

// ObserveResolverError records a GraphQL resolver error.
func (r *NoOp) ObserveResolverError(string, string, string) {}
