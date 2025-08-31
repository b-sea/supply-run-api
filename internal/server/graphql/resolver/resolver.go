package resolver

import "github.com/b-sea/supply-run-api/internal/auth"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Auth *auth.Service
}
