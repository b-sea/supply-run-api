// Package metrics implements all metric recorders.
package metrics

import (
	"time"

	"github.com/b-sea/go-server/metrics"
	"github.com/b-sea/go-server/server"
	"github.com/b-sea/supply-run-api/internal/graphql"
)

var (
	_ server.Recorder  = (*NoOp)(nil)
	_ graphql.Recorder = (*NoOp)(nil)
)

// NoOp is a simple metrics recorder that does nothing.
type NoOp struct {
	metrics.NoOp
}

// NewNoOp creates a new NoOp.
func NewNoOp() *NoOp {
	return &NoOp{}
}

// ObserveResolverDuration records the duration of a GraphQL resolver.
func (r *NoOp) ObserveResolverDuration(string, string, string, time.Duration) {}

// ObserveGraphqlError records an unhandled GraphQL error.
func (r *NoOp) ObserveGraphqlError() {}
