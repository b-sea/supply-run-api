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
	"github.com/sirupsen/logrus"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func NewHandler(authService *auth.Service) http.Handler {
	schema := resolver.NewExecutableSchema(
		resolver.Config{
			Resolvers: &resolver.Resolver{
				Auth: authService,
			},
			Directives: resolver.DirectiveRoot{
				Protected: authService.Directive(),
			},
		},
	)
	server := handler.NewDefaultServer(schema)

	server.AroundOperations(schemaIntrospection(false))
	server.AroundFields(cleanFieldErrorMiddleware())
	server.AroundFields(metricsMiddleware())

	return server
}

func schemaIntrospection(disable bool) func(context.Context, graphql.OperationHandler) graphql.ResponseHandler {
	return func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
		graphql.GetOperationContext(ctx).DisableIntrospection = disable
		return next(ctx)
	}
}

func metricsMiddleware() func(context.Context, graphql.Resolver) (interface{}, error) {
	return func(ctx context.Context, next graphql.Resolver) (interface{}, error) {
		start := time.Now()
		result, err := next(ctx)

		field := graphql.GetFieldContext(ctx)
		if !field.IsResolver {
			return result, err
		}

		if err != nil {
			logrus.Infof("%s.%s, %s", field.Object, field.Field.Name, asGraphQLError(err).Extensions["code"])
		}

		logrus.Infof("%s.%s: %s", field.Object, field.Field.Name, time.Since(start))

		return result, err
	}
}

func cleanFieldErrorMiddleware() func(context.Context, graphql.Resolver) (interface{}, error) {
	return func(ctx context.Context, next graphql.Resolver) (interface{}, error) {
		result, err := next(ctx)
		if err != nil {
			return nil, asGraphQLError(err)
		}

		return result, nil
	}
}

func asGraphQLError(err error) *gqlerror.Error {
	var gqlErr *gqlerror.Error
	if !errors.As(err, &gqlErr) {
		gqlErr = gqlerror.Wrap(err)
	}

	errcode.Set(gqlErr, "ERR_"+strings.ToUpper(strings.Replace(gqlErr.Message, " ", "_", -1)))
	return gqlErr
}
