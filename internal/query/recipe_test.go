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
		repo   query.RecipeRepository
		filter query.RecipeFilter
		page   query.Pagination
		order  query.Order
		result *query.RecipePage
		err    error
	}

	tests := map[string]testCase{
		"first page": {
			repo: &mock.QueryRecipeRepository{
				FindRecipesResult: []*query.Recipe{
					{ID: entity.NewID("1")},
					{ID: entity.NewID("2")},
					{ID: entity.NewID("3")},
					{ID: entity.NewID("4")},
					{ID: entity.NewID("5")},
				},
				FindRecipesErr: nil,
			},
			filter: query.RecipeFilter{},
			page: query.Pagination{
				Size: 3,
			},
			order: query.Order{},
			result: &query.RecipePage{
				Info: query.PageInfo{
					HasNextPage:     true,
					HasPreviousPage: false,
					StartCursor: &query.Cursor{
						ID:   entity.NewID("1"),
						Sort: query.CreatedSort,
					},
					EndCursor: &query.Cursor{
						ID:   entity.NewID("3"),
						Sort: query.CreatedSort,
					},
				},
				Items: []*query.Recipe{
					{ID: entity.NewID("1")},
					{ID: entity.NewID("2")},
					{ID: entity.NewID("3")},
				},
			},
			err: nil,
		},
		"middle page": {
			repo: &mock.QueryRecipeRepository{
				FindRecipesResult: []*query.Recipe{
					{ID: entity.NewID("2")},
					{ID: entity.NewID("3")},
					{ID: entity.NewID("4")},
					{ID: entity.NewID("5")},
					{ID: entity.NewID("6")},
				},
				FindRecipesErr: nil,
			},
			filter: query.RecipeFilter{},
			page: query.Pagination{
				Cursor: &query.Cursor{ID: entity.NewID("2")},
				Size:   3,
			},
			order: query.Order{},
			result: &query.RecipePage{
				Info: query.PageInfo{
					HasNextPage:     true,
					HasPreviousPage: true,
					StartCursor: &query.Cursor{
						ID:   entity.NewID("3"),
						Sort: query.CreatedSort,
					},
					EndCursor: &query.Cursor{
						ID:   entity.NewID("5"),
						Sort: query.CreatedSort,
					},
				},
				Items: []*query.Recipe{
					{ID: entity.NewID("3")},
					{ID: entity.NewID("4")},
					{ID: entity.NewID("5")},
				},
			},
			err: nil,
		},
		"last page": {
			repo: &mock.QueryRecipeRepository{
				FindRecipesResult: []*query.Recipe{
					{ID: entity.NewID("5")},
					{ID: entity.NewID("6")},
					{ID: entity.NewID("7")},
				},
				FindRecipesErr: nil,
			},
			filter: query.RecipeFilter{},
			page: query.Pagination{
				Cursor: &query.Cursor{ID: entity.NewID("5")},
				Size:   3,
			},
			order: query.Order{},
			result: &query.RecipePage{
				Info: query.PageInfo{
					HasNextPage:     false,
					HasPreviousPage: true,
					StartCursor: &query.Cursor{
						ID:   entity.NewID("6"),
						Sort: query.CreatedSort,
					},
					EndCursor: &query.Cursor{
						ID:   entity.NewID("7"),
						Sort: query.CreatedSort,
					},
				},
				Items: []*query.Recipe{
					{ID: entity.NewID("6")},
					{ID: entity.NewID("7")},
				},
			},
			err: nil,
		},
		"no page size": {
			repo: &mock.QueryRecipeRepository{
				FindRecipesResult: []*query.Recipe{},
				FindRecipesErr:    nil,
			},
			filter: query.RecipeFilter{},
			page:   query.Pagination{},
			order:  query.Order{},
			result: &query.RecipePage{
				Info:  query.PageInfo{},
				Items: []*query.Recipe{},
			},
			err: nil,
		},
		"no find results": {
			repo: &mock.QueryRecipeRepository{
				FindRecipesResult: []*query.Recipe{},
				FindRecipesErr:    nil,
			},
			filter: query.RecipeFilter{},
			page: query.Pagination{
				Size: 3,
			},
			order: query.Order{},
			result: &query.RecipePage{
				Info:  query.PageInfo{},
				Items: []*query.Recipe{},
			},
			err: nil,
		},
		"unknown error": {
			repo: &mock.QueryRecipeRepository{
				FindRecipesResult: nil,
				FindRecipesErr:    errors.New("something went wrong"),
			},
			filter: query.RecipeFilter{},
			page: query.Pagination{
				Size: 3,
			},
			order:  query.Order{},
			result: nil,
			err:    errors.New("something went wrong"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			service := query.NewService(test.repo, &mock.QueryUnitRepository{}, &mock.QueryUserRepository{})
			result, err := service.FindRecipes(context.Background(), test.filter, test.page, test.order)

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
		repo   query.RecipeRepository
		id     entity.ID
		result *query.Recipe
		err    error
	}

	tests := map[string]testCase{
		"success": {
			repo: &mock.QueryRecipeRepository{
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
			repo: &mock.QueryRecipeRepository{
				GetRecipesResult: []*query.Recipe{},
				GetRecipesErr:    nil,
			},
			id:     entity.NewID("recipe-123"),
			result: nil,
			err:    entity.ErrNotFound,
		},
		"unknown error": {
			repo: &mock.QueryRecipeRepository{
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
			service := query.NewService(test.repo, &mock.QueryUnitRepository{}, &mock.QueryUserRepository{})
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

func TestFindTags(t *testing.T) {
	t.Parallel()

	type testCase struct {
		repo   query.RecipeRepository
		result []string
		err    error
	}

	tests := map[string]testCase{
		"success": {
			repo: &mock.QueryRecipeRepository{
				FindTagsResult: []string{"vegan", "carnivore"},
				FindTagsErr:    nil,
			},
			result: []string{
				"carnivore",
				"vegan",
			},
			err: nil,
		},
		"unknown error": {
			repo: &mock.QueryRecipeRepository{
				FindTagsResult: nil,
				FindTagsErr:    errors.New("something went wrong"),
			},
			result: nil,
			err:    errors.New("something went wrong"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			service := query.NewService(test.repo, &mock.QueryUnitRepository{}, &mock.QueryUserRepository{})
			result, err := service.FindTags(context.Background(), nil)

			assert.Equal(t, test.result, result)
			if test.err == nil {
				assert.NoError(t, err)
			} else {
				assert.ErrorAs(t, err, &test.err)
			}
		})
	}
}
