// Package resolver implements all GraphQL resolvers.
package resolver

import "github.com/b-sea/supply-run-api/internal/auth"

// Resolver defines all data available to the resolvers.
type Resolver struct {
	Auth *auth.Service
}
