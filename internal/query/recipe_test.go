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

func TestFindRecipe(t *testing.T) {
	t.Parallel()

	type testCase struct {
		repo   query.Repository
		filter *query.RecipeFilter
		page   *query.Pagination
		result []*query.Recipe
		err    error
	}

	tests := map[string]testCase{
		"nil inputs": {
			repo: &mock.QueryRepository{
				FindRecipesResult: []*query.Recipe{
					{ID: entity.NewID("recipe-123")},
				},
				FindRecipesErr: nil,
			},
			filter: nil,
			page:   nil,
			result: []*query.Recipe{
				{ID: entity.NewID("recipe-123")},
			},
			err: nil,
		},
		"unknown error": {
			repo: &mock.QueryRepository{
				FindRecipesResult: nil,
				FindRecipesErr:    errors.New("something went wrong"),
			},
			filter: nil,
			page:   nil,
			result: nil,
			err:    errors.New("something went wrong"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			service := query.NewService(test.repo)
			result, err := service.FindRecipes(context.Background(), test.filter, test.page)

			assert.Equal(t, test.result, result)
			if test.err == nil {
				assert.NoError(t, err)
			} else {
				assert.ErrorAs(t, err, &test.err)
			}
		})
	}
}

func TestGetRecipe(t *testing.T) {
	t.Parallel()

	type testCase struct {
		repo   query.Repository
		id     entity.ID
		result *query.Recipe
		err    error
	}

	tests := map[string]testCase{
		"success": {
			repo: &mock.QueryRepository{
				GetRecipesResult: []*query.Recipe{
					{ID: entity.NewID("recipe-123")},
				},
				GetRecipesErr: nil,
			},
			id: entity.NewID("recipe-123"),
			result: &query.Recipe{
				ID: entity.NewID("recipe-123"),
			},
			err: nil,
		},
		"not found": {
			repo: &mock.QueryRepository{
				GetRecipesResult: []*query.Recipe{},
				GetRecipesErr:    nil,
			},
			id:     entity.NewID("recipe-123"),
			result: nil,
			err:    entity.ErrNotFound,
		},
		"unknown error": {
			repo: &mock.QueryRepository{
				GetRecipesResult: nil,
				GetRecipesErr:    errors.New("something went wrong"),
			},
			id:     entity.NewID("recipe-123"),
			result: nil,
			err:    errors.New("something went wrong"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			service := query.NewService(test.repo)
			result, err := service.GetRecipe(context.Background(), test.id)

			assert.Equal(t, test.result, result)
			if test.err == nil {
				assert.NoError(t, err)
			} else {
				assert.ErrorAs(t, err, &test.err)
			}
		})
	}
}

func TestGetIngredients(t *testing.T) {
	t.Parallel()

	type testCase struct {
		repo   query.Repository
		result []string
		err    error
	}

	tests := map[string]testCase{
		"success": {
			repo: &mock.QueryRepository{
				GetIngredientsResult: []string{
					"milk", "bread",
				},
				GetIngredientsErr: nil,
			},
			result: []string{
				"bread", "milk",
			},
			err: nil,
		},
		"unknown error": {
			repo: &mock.QueryRepository{
				GetIngredientsResult: nil,
				GetIngredientsErr:    errors.New("something went wrong"),
			},
			result: nil,
			err:    errors.New("something went wrong"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			service := query.NewService(test.repo)
			result, err := service.GetIngredients(context.Background())

			assert.Equal(t, test.result, result)
			if test.err == nil {
				assert.NoError(t, err)
			} else {
				assert.ErrorAs(t, err, &test.err)
			}
		})
	}
}

func TestGetTags(t *testing.T) {
	t.Parallel()

	type testCase struct {
		repo   query.Repository
		result []string
		err    error
	}

	tests := map[string]testCase{
		"success": {
			repo: &mock.QueryRepository{
				GetTagsResult: []string{
					"gluten free", "breakfast",
				},
				GetTagsErr: nil,
			},
			result: []string{
				"breakfast", "gluten free",
			},
			err: nil,
		},
		"unknown error": {
			repo: &mock.QueryRepository{
				GetTagsResult: nil,
				GetTagsErr:    errors.New("something went wrong"),
			},
			result: nil,
			err:    errors.New("something went wrong"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			service := query.NewService(test.repo)
			result, err := service.GetTags(context.Background())

			assert.Equal(t, test.result, result)
			if test.err == nil {
				assert.NoError(t, err)
			} else {
				assert.ErrorAs(t, err, &test.err)
			}
		})
	}
}
