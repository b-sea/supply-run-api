// Package metrics implements all metric recorders.
package metrics

import (
	"time"

	"github.com/b-sea/go-server/metrics"
	"github.com/b-sea/go-server/server"
	"github.com/b-sea/supply-run-api/internal/graphql"
	"github.com/b-sea/supply-run-api/internal/mariadb"
)

var (
	_ server.Recorder  = (*NoOp)(nil)
	_ graphql.Recorder = (*NoOp)(nil)
	_ mariadb.Recorder = (*NoOp)(nil)
)

// NoOp is a simple metrics recorder that does nothing.
type NoOp struct {
	metrics.NoOp
}

// NewNoOp creates a new NoOp.
func NewNoOp() *NoOp {
	return &NoOp{}
}

// ObserveGraphqlResolverDuration records the duration of a GraphQL resolver.
func (r *NoOp) ObserveGraphqlResolverDuration(string, string, string, time.Duration) {}

// ObserveGraphqlError records an unhandled GraphQL error.
func (r *NoOp) ObserveGraphqlError() {}

// ObserveMariadbTxDuration records the duration of a MariaDB transaction.
func (r *NoOp) ObserveMariadbTxDuration(string, time.Duration) {}
