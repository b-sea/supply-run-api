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
	ObserveResolverDuration(object string, field string, status string, duration time.Duration)
	ObserveGraphqlError()
}

func operationTelemetry() graphql.OperationMiddleware {
	return func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
		op := graphql.GetOperationContext(ctx)

		log := zerolog.Ctx(ctx)
		log.UpdateContext(func(c zerolog.Context) zerolog.Context {
			return c.Str("operation", string(op.Operation.Operation))
		})

		return next(log.WithContext(ctx))
	}
}

func fieldTelemetry(recorder Recorder) graphql.FieldMiddleware {
	return func(ctx context.Context, next graphql.Resolver) (any, error) {
		start := time.Now()

		// Only log resolvers data
		field := graphql.GetFieldContext(ctx)
		if !field.IsResolver {
			return next(ctx)
		}

		log := zerolog.Ctx(ctx)
		log.UpdateContext(func(c zerolog.Context) zerolog.Context {
			return c.Str("object", field.Object).Str("field", field.Field.Name)
		})

		result, err := next(log.WithContext(ctx))

		defer func() {
			event := log.Info() //nolint: zerologlint
			status := "success"

			if err != nil {
				status = "failed"
				event = event.Err(err)
			}

			duration := time.Since(start)

			event.Str("status", status).Dur("duration_ms", duration).Msg("resolver complete")
			recorder.ObserveResolverDuration(field.Object, field.Field.Name, status, duration)
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
