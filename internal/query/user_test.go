package query_test

import (
	"context"
	"errors"
	"testing"

	"github.com/b-sea/supply-run-api/internal/entity"
	"github.com/b-sea/supply-run-api/internal/mock"
	"github.com/b-sea/supply-run-api/internal/query"
	"github.com/stretchr/testify/assert"
)

func TestGetUsers(t *testing.T) {
	t.Parallel()

	type testCase struct {
		repo   query.Repository
		ids    []entity.ID
		result []*query.User
		err    error
	}

	tests := map[string]testCase{
		"success": {
			repo: &mock.QueryRepository{
				GetUsersResult: []*query.User{
					{ID: entity.NewID("recipe-123")},
				},
				GetUsersErr: nil,
			},
			ids: []entity.ID{entity.NewID("recipe-123")},
			result: []*query.User{
				{ID: entity.NewID("recipe-123")},
			},
			err: nil,
		},
		"unknown error": {
			repo: &mock.QueryRepository{
				GetUsersResult: nil,
				GetUsersErr:    errors.New("something went wrong"),
			},
			ids:    []entity.ID{entity.NewID("recipe-123")},
			result: nil,
			err:    errors.New("something went wrong"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			service := query.NewService(test.repo)
			result, err := service.GetUsers(context.Background(), test.ids)

			assert.Equal(t, test.result, result)
			if test.err == nil {
				assert.NoError(t, err)
			} else {
				assert.ErrorAs(t, err, &test.err)
			}
		})
	}
}
