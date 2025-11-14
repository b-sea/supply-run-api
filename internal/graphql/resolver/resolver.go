// Package resolver implements all GraphQL resolvers.
package resolver

import "github.com/b-sea/supply-run-api/internal/query"

// Resolver defines all data available to the resolvers.
type Resolver struct {
	queries *query.Service
}

// NewResolver creates a new Resolver.
func NewResolver(queries *query.Service) *Resolver {
	return &Resolver{
		queries: queries,
	}
}
