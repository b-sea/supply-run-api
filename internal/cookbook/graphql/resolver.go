package graphql

import (
	"github.com/b-sea/supply-run-api/internal/auth"
	"github.com/b-sea/supply-run-api/internal/cookbook/app"
)

type Resolver struct {
	Auth *auth.Service
	App  *app.Application
}
