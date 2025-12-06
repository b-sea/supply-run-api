package metrics

import (
	"time"

	"github.com/b-sea/go-server/metrics"
	"github.com/b-sea/go-server/server"
	"github.com/b-sea/supply-run-api/internal/graphql"
	"github.com/b-sea/supply-run-api/internal/mariadb"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	namespace        = "supply_run"
	graphqlSubsystem = "graphql"
	mariadbSubsystem = "mariadb"
)

var (
	_ server.Recorder  = (*Prometheus)(nil)
	_ graphql.Recorder = (*Prometheus)(nil)
	_ mariadb.Recorder = (*Prometheus)(nil)
)

// Prometheus is a metrics recorder for Prometheus.
type Prometheus struct {
	metrics.Prometheus

	resolverDuration  *prometheus.HistogramVec
	graphqlError      prometheus.Counter
	mariaDBTxDuration *prometheus.HistogramVec
}

// NewPrometheus creates a new Prometheus recorder.
func NewPrometheus() *Prometheus {
	recorder := &Prometheus{
		Prometheus: *metrics.NewPrometheus(namespace, metrics.WithGroupedCodes()),
		resolverDuration: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Namespace: namespace,
				Subsystem: graphqlSubsystem,
				Name:      "resolver_duration",
				Help:      "GraphQL Resolver Duration in Seconds",
			},
			[]string{"object", "field", "status"},
		),
		graphqlError: prometheus.NewCounter(
			prometheus.CounterOpts{
				Namespace: namespace,
				Subsystem: graphqlSubsystem,
				Name:      "error_total",
				Help:      "Unhandled GraphQL Errors",
			},
		),
		mariaDBTxDuration: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Namespace: namespace,
				Subsystem: mariadbSubsystem,
				Name:      "transaction_duration",
				Help:      "MariaDB Transaction Duration in Seconds",
			},
			[]string{"status"},
		),
	}

	_ = prometheus.DefaultRegisterer.Register(recorder.resolverDuration)
	_ = prometheus.DefaultRegisterer.Register(recorder.graphqlError)

	return recorder
}

// ObserveResolverDuration records the duration of a GraphQL resolver.
func (p *Prometheus) ObserveResolverDuration(object string, field string, status string, duration time.Duration) {
	p.resolverDuration.WithLabelValues(object, field, status).Observe(duration.Seconds())
}

// ObserveGraphqlError records an unhandled GraphQL error.
func (p *Prometheus) ObserveGraphqlError() {
	p.graphqlError.Inc()
}

// ObserveMariaDBTxDuration records the duration of a MariaDB transaction.
func (p *Prometheus) ObserveMariaDBTxDuration(status string, duration time.Duration) {
	p.mariaDBTxDuration.WithLabelValues(status).Observe(duration.Seconds())
}
