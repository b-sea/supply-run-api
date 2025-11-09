// Package graphql implements a GraphQL API.
package graphql

import (
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/b-sea/supply-run-api/internal/server/graphql/resolver"
)

// Recorder defines functions for tracking GraphQL-based metrics.
type Recorder interface {
	ObserveResolverDuration(object string, field string, status string, duration time.Duration)
	ObserveResolverError(object string, field string, code string)
}

// NewHandler configures and creates a GraphQL API handler.
func NewHandler(recorder Recorder) http.Handler { // coverage-ignore
	schema := resolver.NewExecutableSchema(
		resolver.Config{Resolvers: resolver.NewResolver()},
	)

	server := handler.NewDefaultServer(schema)
	server.AroundFields(fieldTelemetry(recorder))

	return server
}
