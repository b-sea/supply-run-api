// Package graphql implements a GraphQL API.
package graphql

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/b-sea/supply-run-api/internal/graphql/resolver"
	"github.com/b-sea/supply-run-api/internal/query"
)

// NewHandler configures and creates a GraphQL API handler.
func NewHandler(queries *query.Service, recorder Recorder) http.Handler {
	schema := resolver.NewExecutableSchema(
		resolver.Config{Resolvers: resolver.NewResolver(queries)},
	)

	server := handler.NewDefaultServer(schema)
	server.AroundFields(fieldTelemetry(recorder))
	server.AroundOperations(operationTelemetry())
	server.SetRecoverFunc(recoverTelemetry(recorder))

	return server
}
