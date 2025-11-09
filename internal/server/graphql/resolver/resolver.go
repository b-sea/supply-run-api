// Package resolver implements all GraphQL resolvers.
package resolver

// Resolver defines all data available to the resolvers.
type Resolver struct{}

// NewResolver creates a new Resolver.
func NewResolver() *Resolver {
	return &Resolver{}
}
