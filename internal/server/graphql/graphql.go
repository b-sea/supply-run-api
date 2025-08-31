// Package graphql implements a GraphQL API.
package graphql

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/errcode"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/b-sea/supply-run-api/internal/auth"
	"github.com/b-sea/supply-run-api/internal/server/graphql/resolver"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// Recorder defines functions for tracking GraphQL-based metrics.
type Recorder interface {
	ObserveResolverDuration(object string, field string, status string, duration time.Duration)
	ObserveResolverError(object string, field string, code string)
}

// NewHandler configures and creates a GraphQL API handler.
func NewHandler(authN *auth.Service, recorder Recorder) http.Handler {
	schema := resolver.NewExecutableSchema(
		resolver.Config{
			Resolvers: &resolver.Resolver{
				Auth: authN,
			},
			Directives: resolver.DirectiveRoot{
				Protected: func(ctx context.Context, _ interface{}, next graphql.Resolver) (interface{}, error) {
					if _, err := auth.FromContext(ctx); err != nil {
						return nil, err
					}

					return next(ctx)
				},
			},
		},
	)
	server := handler.NewDefaultServer(schema)

	server.AroundFields(
		func(ctx context.Context, next graphql.Resolver) (interface{}, error) {
			result, err := next(ctx)
			if err != nil {
				return nil, asGraphQLError(err)
			}

			return result, nil
		},
	)
	server.AroundFields(
		func(ctx context.Context, next graphql.Resolver) (interface{}, error) {
			start := time.Now()
			result, err := next(ctx)

			field := graphql.GetFieldContext(ctx)
			if !field.IsResolver {
				return result, err
			}

			status := "success"

			if err != nil {
				status = "failed"
				code, _ := asGraphQLError(err).Extensions["code"].(string)
				recorder.ObserveResolverError(field.Object, field.Field.Name, code)
			}

			recorder.ObserveResolverDuration(field.Object, field.Field.Name, status, time.Since(start))

			return result, err
		},
	)

	return server
}

func asGraphQLError(err error) *gqlerror.Error {
	var gqlErr *gqlerror.Error
	if !errors.As(err, &gqlErr) {
		gqlErr = gqlerror.Wrap(err)
	}

	errcode.Set(gqlErr, "ERR_"+strings.ToUpper(strings.ReplaceAll(gqlErr.Message, " ", "_")))

	return gqlErr
}
