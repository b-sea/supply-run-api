// Package prometheus implements all interactions with Prometheus.
package prometheus

import (
	"net/http"
	"time"

	"github.com/b-sea/supply-run-api/internal/auth"
	"github.com/b-sea/supply-run-api/internal/server"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	httpSubsystem    = "http"
	graphqlSubsystem = "graphql"
)

var (
	_ server.Recorder = (*Recorder)(nil)
	_ auth.Recorder   = (*Recorder)(nil)
)

// Recorder implements metric recording with Prometheus.
type Recorder struct {
	requestDuration   *prometheus.HistogramVec
	responseSize      *prometheus.HistogramVec
	requestAuthorized *prometheus.CounterVec

	resolverDuration *prometheus.HistogramVec
	resolverError    *prometheus.CounterVec
}

// NewRecorder creates a new Prometheus recorder.
func NewRecorder(namespace string) *Recorder {
	recorder := &Recorder{
		requestDuration: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Namespace: namespace,
				Subsystem: httpSubsystem,
				Name:      "request_duration_seconds",
				Help:      "HTTP request latency",
			},
			[]string{"method", "path", "code"},
		),
		responseSize: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Namespace: namespace,
				Subsystem: httpSubsystem,
				Name:      "response_size_bytes",
				Help:      "HTTP response size (in bytes)",
				Buckets:   prometheus.ExponentialBucketsRange(10, 5e+7, 10), //nolint: mnd
			},
			[]string{"method", "path", "code"},
		),
		requestAuthorized: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: namespace,
				Subsystem: httpSubsystem,
				Name:      "authorized_request_total",
				Help:      "Number of authorized requests per user",
			},
			[]string{"user"},
		),
		resolverDuration: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Namespace: namespace,
				Subsystem: graphqlSubsystem,
				Name:      "resolver_duration_seconds",
				Help:      "GraphQL resolver latency",
			},
			[]string{"object", "field", "status"},
		),
		resolverError: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: namespace,
				Subsystem: graphqlSubsystem,
				Name:      "resolver_error",
				Help:      "GraphQL resolver error",
			},
			[]string{"object", "field", "code"},
		),
	}

	prometheus.DefaultRegisterer.MustRegister(recorder.requestDuration)
	prometheus.DefaultRegisterer.MustRegister(recorder.responseSize)
	prometheus.DefaultRegisterer.MustRegister(recorder.requestAuthorized)

	prometheus.DefaultRegisterer.MustRegister(recorder.resolverDuration)
	prometheus.DefaultRegisterer.MustRegister(recorder.resolverError)

	return recorder
}

// Handler returns an HTTP handler for Prometheus.
func (r *Recorder) Handler() http.Handler {
	return promhttp.Handler()
}

// ObserveRequestDuration updates the HTTP request duration metric.
func (r *Recorder) ObserveRequestDuration(method string, path string, code string, duration time.Duration) {
	r.requestDuration.WithLabelValues(method, path, code).Observe(duration.Seconds())
}

// ObserveResponseSize updates the HTTP response size metric.
func (r *Recorder) ObserveResponseSize(method string, path string, code string, bytes int64) {
	r.responseSize.WithLabelValues(method, path, code).Observe(float64(bytes))
}

// RequestAuthorized records that a known user made a request.
func (r *Recorder) RequestAuthorized(username string) {
	r.requestAuthorized.WithLabelValues(username).Inc()
}

// ObserveResolverDuration updates the GraphQL resolver duration metric.
func (r *Recorder) ObserveResolverDuration(object string, field string, status string, duration time.Duration) {
	r.resolverDuration.WithLabelValues(object, field, status).Observe(duration.Seconds())
}

// ObserveResolverError records that a GraphQL resolver has errored.
func (r *Recorder) ObserveResolverError(object string, field string, code string) {
	r.resolverError.WithLabelValues(object, field, code).Inc()
}
