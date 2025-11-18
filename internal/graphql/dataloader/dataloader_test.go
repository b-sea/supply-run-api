package dataloader_test

import (
	"context"
	"errors"
	"testing"

	"github.com/b-sea/supply-run-api/internal/entity"
	"github.com/b-sea/supply-run-api/internal/graphql/dataloader"
	"github.com/b-sea/supply-run-api/internal/graphql/model"
	"github.com/b-sea/supply-run-api/internal/mock"
	"github.com/b-sea/supply-run-api/internal/query"
	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	t.Parallel()

	type testCase struct {
		ctx    context.Context
		id     entity.ID
		result model.UserResult
		err    error
	}

	tests := map[string]testCase{
		"success": {
			ctx: dataloader.ToContext(
				context.Background(),
				dataloader.New(
					query.NewService(
						&mock.QueryRepository{
							GetUsersResult: []*query.User{{ID: entity.NewID("1234")}},
						},
					),
				),
			),
			id:     entity.NewID("1234"),
			result: &model.User{ID: model.NewUserID(entity.NewID("1234"))},
			err:    nil,
		},
		"empty context": {
			ctx:    context.Background(),
			id:     entity.NewID("1234"),
			result: nil,
			err:    dataloader.ErrDataloader,
		},
		"repo error": {
			ctx: dataloader.ToContext(
				context.Background(),
				dataloader.New(
					query.NewService(
						&mock.QueryRepository{
							GetUsersErr: errors.New("something went wrong"),
						},
					),
				),
			),
			id:     entity.NewID("1234"),
			result: nil,
			err:    errors.New("something went wrong"),
		},
		"mixed results": {
			ctx: dataloader.ToContext(
				context.Background(),
				dataloader.New(
					query.NewService(
						&mock.QueryRepository{
							GetUsersResult: []*query.User{{ID: entity.NewID("9999")}},
						},
					),
				),
			),
			id:     entity.NewID("1234"),
			result: &model.NotFoundError{ID: model.NewUserID(entity.NewID("1234"))},
			err:    nil,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result, err := dataloader.GetUser(test.ctx, test.id)

			assert.Equal(t, test.result, result)
			if test.err == nil {
				assert.NoError(t, err)
			} else {
				assert.ErrorAs(t, err, &test.err)
			}
		})
	}
}
