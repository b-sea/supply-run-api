package graphql

import (
	"context"
	"fmt"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// Recorder defines functions for tracking GraphQL-based metrics.
type Recorder interface {
	ObserveGraphqlResolverDuration(object string, field string, status string, duration time.Duration)
	ObserveGraphqlError()
}

func fieldTelemetry(recorder Recorder) graphql.FieldMiddleware {
	return func(ctx context.Context, next graphql.Resolver) (any, error) {
		start := time.Now()

		// Only log resolvers data
		field := graphql.GetFieldContext(ctx)
		if !field.IsResolver {
			return next(ctx)
		}

		result, err := next(ctx)

		defer func() {
			event := zerolog.Ctx(ctx).Info() //nolint: zerologlint
			status := "success"

			if err != nil {
				status = "failed"
				event = event.Err(err)
			}

			duration := time.Since(start)

			event.
				Str("object", field.Object).
				Str("field", field.Field.Name).
				Str("status", status).
				Dur("duration_ms", duration).
				Msg("resolver complete")
			recorder.ObserveGraphqlResolverDuration(field.Object, field.Field.Name, status, duration)
		}()

		return result, err
	}
}

func recoverTelemetry(recorder Recorder) graphql.RecoverFunc {
	return func(ctx context.Context, err any) error {
		asErr, ok := err.(error)
		if !ok {
			asErr = fmt.Errorf("%v", err) //nolint: err113
		}

		zerolog.Ctx(ctx).Error().Stack().Err(errors.Wrap(asErr, "graphql")).Msg("unhandled error")
		recorder.ObserveGraphqlError()

		return gqlerror.Errorf("internal server error")
	}
}
