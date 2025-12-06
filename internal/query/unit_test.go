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

func TestGetUnits(t *testing.T) {
	t.Parallel()

	type testCase struct {
		repo   query.UnitRepository
		ids    []entity.ID
		result []*query.Unit
		err    error
	}

	tests := map[string]testCase{
		"success": {
			repo: &mock.QueryUnitRepository{
				GetUnitsResult: []*query.Unit{
					{ID: entity.NewID("unit-123")},
				},
				GetUnitsErr: nil,
			},
			ids: []entity.ID{entity.NewID("unit-123")},
			result: []*query.Unit{
				{ID: entity.NewID("unit-123")},
			},
			err: nil,
		},
		"unknown error": {
			repo: &mock.QueryUnitRepository{
				GetUnitsResult: nil,
				GetUnitsErr:    errors.New("something went wrong"),
			},
			ids:    []entity.ID{entity.NewID("unit-123")},
			result: nil,
			err:    errors.New("something went wrong"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			service := query.NewService(&mock.QueryRecipeRepository{}, test.repo, &mock.QueryUserRepository{})
			result, err := service.GetUnits(context.Background(), test.ids)

			assert.Equal(t, test.result, result)
			if test.err == nil {
				assert.NoError(t, err)
			} else {
				assert.ErrorAs(t, err, &test.err)
			}
		})
	}
}
