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

	graphqlResolverDuration *prometheus.HistogramVec
	graphqlError            prometheus.Counter
	mariadbTxDuration       *prometheus.HistogramVec
}

// NewPrometheus creates a new Prometheus recorder.
func NewPrometheus() *Prometheus {
	recorder := &Prometheus{
		Prometheus: *metrics.NewPrometheus(namespace, metrics.WithGroupedCodes()),
		graphqlResolverDuration: prometheus.NewHistogramVec(
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
		mariadbTxDuration: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Namespace: namespace,
				Subsystem: mariadbSubsystem,
				Name:      "transaction_duration",
				Help:      "MariaDB Transaction Duration in Seconds",
			},
			[]string{"status"},
		),
	}

	_ = prometheus.DefaultRegisterer.Register(recorder.graphqlResolverDuration)
	_ = prometheus.DefaultRegisterer.Register(recorder.graphqlError)
	_ = prometheus.DefaultRegisterer.Register(recorder.mariadbTxDuration)

	return recorder
}

// ObserveGraphqlResolverDuration records the duration of a GraphQL resolver.
func (p *Prometheus) ObserveGraphqlResolverDuration(object string, field string, status string, duration time.Duration) {
	p.graphqlResolverDuration.WithLabelValues(object, field, status).Observe(duration.Seconds())
}

// ObserveGraphqlError records an unhandled GraphQL error.
func (p *Prometheus) ObserveGraphqlError() {
	p.graphqlError.Inc()
}

// ObserveMariaDBTxDuration records the duration of a MariaDB transaction.
func (p *Prometheus) ObserveMariadbTxDuration(status string, duration time.Duration) {
	p.mariadbTxDuration.WithLabelValues(status).Observe(duration.Seconds())
}
