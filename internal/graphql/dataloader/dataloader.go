// Package dataloader implements GraphQL dataloaders to solve the N+1 problem.
package dataloader

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/b-sea/supply-run-api/internal/entity"
	"github.com/b-sea/supply-run-api/internal/graphql/model"
	"github.com/b-sea/supply-run-api/internal/query"
	"github.com/graph-gophers/dataloader"
	"github.com/rs/zerolog"
)

type ctxKey string

const loaderKey = ctxKey("dataloader")

// ErrDataloader is thrown when the dataloader has an unexpected error.
var ErrDataloader = errors.New("dataloader error")

func dataloaderError(v any) error {
	return fmt.Errorf("%w: %v", ErrDataloader, v)
}

// Dataloader batches and consolidates data calls.
type Dataloader struct {
	getUnit *dataloader.Loader
	getUser *dataloader.Loader
}

// New creates a new Dataloader.
func New(queries *query.Service) *Dataloader {
	return &Dataloader{
		getUnit: dataloader.NewBatchedLoader(batchGetUnit(queries)),
		getUser: dataloader.NewBatchedLoader(batchGetUser(queries)),
	}
}

// GetUnit returns a UnitResult from an ID.
func GetUnit(ctx context.Context, id entity.ID) (model.UnitResult, error) { //nolint: ireturn
	loader, err := FromContext(ctx)
	if err != nil {
		return nil, err
	}

	data, err := loader.getUnit.Load(ctx, dataloader.StringKey(id.String()))()
	if err != nil {
		return nil, err
	}

	result, _ := data.(model.UnitResult)

	return result, nil
}

func batchGetUnit(queries *query.Service) dataloader.BatchFunc { //nolint: dupl
	return func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		start := time.Now()

		defer func() {
			zerolog.Ctx(ctx).Info().
				Dur("duration_ms", time.Since(start)).
				Int("batch", len(keys)).
				Msg("unit dataloader complete")
		}()

		keyOrder := make(map[entity.ID]int, len(keys))
		ids := make([]entity.ID, len(keys))

		for i, key := range keys {
			ids[i] = entity.NewID(key.String())
			keyOrder[ids[i]] = i
		}

		results := make([]*dataloader.Result, len(keys))

		units, err := queries.GetUnits(ctx, ids)
		if err != nil {
			for i := range keys {
				results[i] = &dataloader.Result{Error: err}
			}

			return results
		}

		for _, unit := range units {
			i, ok := keyOrder[unit.ID]
			if !ok {
				results[i] = &dataloader.Result{
					Data: &model.NotFoundError{ID: model.NewUnitID(unit.ID)},
				}

				continue
			}

			results[i] = &dataloader.Result{
				Data: model.NewUnit(unit),
			}

			delete(keyOrder, unit.ID)
		}

		for id, i := range keyOrder {
			results[i] = &dataloader.Result{
				Data: &model.NotFoundError{ID: model.NewUnitID(id)},
			}
		}

		return results
	}
}

// GetUser returns a UserResult from an ID.
func GetUser(ctx context.Context, id entity.ID) (model.UserResult, error) { //nolint: ireturn
	loader, err := FromContext(ctx)
	if err != nil {
		return nil, err
	}

	data, err := loader.getUser.Load(ctx, dataloader.StringKey(id.String()))()
	if err != nil {
		return nil, err
	}

	result, _ := data.(model.UserResult)

	return result, nil
}

func batchGetUser(queries *query.Service) dataloader.BatchFunc { //nolint: dupl
	return func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		start := time.Now()

		defer func() {
			zerolog.Ctx(ctx).Info().
				Dur("duration_ms", time.Since(start)).
				Int("batch", len(keys)).
				Msg("user dataloader complete")
		}()

		keyOrder := make(map[entity.ID]int, len(keys))
		ids := make([]entity.ID, len(keys))

		for i, key := range keys {
			ids[i] = entity.NewID(key.String())
			keyOrder[ids[i]] = i
		}

		results := make([]*dataloader.Result, len(keys))

		users, err := queries.GetUsers(ctx, ids)
		if err != nil {
			for i := range keys {
				results[i] = &dataloader.Result{Error: err}
			}

			return results
		}

		for _, user := range users {
			i, ok := keyOrder[user.ID]
			if !ok {
				results[i] = &dataloader.Result{
					Data: &model.NotFoundError{ID: model.NewUserID(user.ID)},
				}

				continue
			}

			results[i] = &dataloader.Result{
				Data: model.NewUser(user),
			}

			delete(keyOrder, user.ID)
		}

		for id, i := range keyOrder {
			results[i] = &dataloader.Result{
				Data: &model.NotFoundError{ID: model.NewUserID(id)},
			}
		}

		return results
	}
}

// FromContext returns a Dataloader if it exists in the given Context.
func FromContext(ctx context.Context) (*Dataloader, error) {
	loader, ok := ctx.Value(loaderKey).(*Dataloader)
	if !ok {
		return nil, dataloaderError("not found in context")
	}

	return loader, nil
}

// ToContext stores a Dataloader in the given Context.
func ToContext(ctx context.Context, loader *Dataloader) context.Context {
	return context.WithValue(ctx, loaderKey, loader)
}

// Middleware ensures that a Dataloader exists in a request Context.
func Middleware(queries *query.Service, next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		request = request.WithContext(ToContext(request.Context(), New(queries)))
		next.ServeHTTP(writer, request)
	})
}
