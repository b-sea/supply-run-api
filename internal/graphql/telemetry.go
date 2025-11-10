package graphql

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/errcode"
	"github.com/b-sea/go-logger/logger"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func fieldTelemetry(recorder Recorder) graphql.FieldMiddleware {
	return func(ctx context.Context, next graphql.Resolver) (any, error) {
		start := time.Now()
		log := logger.FromContext(ctx)
		result, err := next(ctx)

		field := graphql.GetFieldContext(ctx)
		if !field.IsResolver {
			return result, err
		}

		defer func() {
			status := "success"

			if err != nil {
				status = "failed"
				code, _ := asGraphQLError(err).Extensions["code"].(string)
				recorder.ObserveResolverError(field.Object, field.Field.Name, code)
				log.Info().
					Str("object", field.Object).
					Str("field", field.Field.Name).
					Str("err_code", code).
					Msg("graphql error")
			}

			duration := time.Since(start)

			log.Info().
				Str("object", field.Object).
				Str("field", field.Field.Name).
				Str("status", status).
				Dur("duration_ms", duration).
				Msg("graphql resolver")

			recorder.ObserveResolverDuration(field.Object, field.Field.Name, status, duration)
		}()

		return result, err
	}
}

func asGraphQLError(err error) *gqlerror.Error { // coverage-ignore
	var gqlErr *gqlerror.Error
	if !errors.As(err, &gqlErr) {
		gqlErr = gqlerror.Wrap(err)
	}

	errcode.Set(gqlErr, "ERR_"+strings.ToUpper(strings.ReplaceAll(gqlErr.Message, " ", "_")))

	return gqlErr
}
