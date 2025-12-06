// Package graphql implements a GraphQL API.
package graphql

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/b-sea/supply-run-api/internal/graphql/dataloader"
	"github.com/b-sea/supply-run-api/internal/graphql/resolver"
	"github.com/b-sea/supply-run-api/internal/query"
)

// GraphQL is an GraphQL API handler.
type GraphQL struct {
	http.Handler
}

// New creates a new GraphQL API handler.
func New(queries *query.Service, recorder Recorder) *GraphQL {
	graphql := &GraphQL{}

	schema := resolver.NewExecutableSchema(
		resolver.Config{
			Resolvers: resolver.NewResolver(queries),
		},
	)

	server := handler.NewDefaultServer(schema)
	server.AroundFields(fieldTelemetry(recorder))
	server.SetRecoverFunc(recoverTelemetry(recorder))

	graphql.Handler = dataloader.Middleware(queries, server)

	return graphql
}
