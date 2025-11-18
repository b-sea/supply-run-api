package model_test

import (
	"testing"
	"time"

	"github.com/b-sea/supply-run-api/internal/entity"
	"github.com/b-sea/supply-run-api/internal/graphql/model"
	"github.com/b-sea/supply-run-api/internal/query"
	"github.com/stretchr/testify/assert"
)

func TestNewRecipeID(t *testing.T) {
	t.Parallel()

	id := entity.NewID("1234")

	result := model.ID{
		Key:  entity.NewID("1234"),
		Kind: model.RecipeKind,
	}

	assert.Equal(t, result, model.NewRecipeID(id))
}

func TestNewRecipe(t *testing.T) {
	t.Parallel()

	created := time.Now().Add(-48 * time.Hour)
	updated := time.Now()

	recipe := &query.Recipe{
		ID:          entity.NewID("1234"),
		Name:        "something good",
		URL:         "test.com",
		NumServings: 45,
		Steps: []string{
			"start", "finish",
		},
		Ingredients: []query.Ingredient{
			{Name: "bread"},
			{Name: "milk"},
		},
		Tags: []string{
			"good", "not good",
		},
		IsFavorite: true,
		CreatedAt:  created,
		CreatedBy:  entity.NewID("creator-123"),
		UpdatedAt:  updated,
		UpdatedBy:  entity.NewID("updater-123"),
	}

	result := &model.Recipe{
		ID:          model.NewRecipeID(entity.NewID("1234")),
		Name:        "something good",
		URL:         "test.com",
		NumServings: 45,
		Steps: []string{
			"start", "finish",
		},
		Ingredients: []*model.Ingredient{
			{Name: "bread"},
			{Name: "milk"},
		},
		Tags: []string{
			"good", "not good",
		},
		IsFavorite:  true,
		CreatedAt:   created,
		CreatedByID: entity.NewID("creator-123"),
		UpdatedAt:   updated,
		UpdatedByID: entity.NewID("updater-123"),
	}

	assert.Equal(t, result, model.NewRecipe(recipe))
}

func TestNewQueryRecipeFilter(t *testing.T) {
	t.Parallel()

	type testCase struct {
		filter *model.RecipeFilter
		result query.RecipeFilter
	}

	name := "something"
	user := model.NewUserID(entity.NewID("user-1234"))
	favorite := true

	tests := map[string]testCase{
		"value": {
			filter: &model.RecipeFilter{
				Name:        &name,
				Ingredients: []string{"bread", "tomato"},
				CreatedBy:   &user,
				IsFavorite:  &favorite,
			},
			result: query.RecipeFilter{
				Name:        &name,
				Ingredients: []string{"bread", "tomato"},
				CreatedBy:   &user.Key,
				IsFavorite:  &favorite,
			},
		},
		"empty": {
			filter: &model.RecipeFilter{},
			result: query.RecipeFilter{},
		},
		"nil": {
			filter: nil,
			result: query.RecipeFilter{},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.result, model.NewQueryRecipeFilter(test.filter))
		})
	}
}

func TestNewRecipeConnection(t *testing.T) {
	t.Parallel()

	type testCase struct {
		page   *query.RecipePage
		result *model.RecipeConnection
	}

	tests := map[string]testCase{
		"default sort": {
			page: &query.RecipePage{
				Info: query.PageInfo{
					HasNextPage:     true,
					HasPreviousPage: true,
					StartCursor:     &query.Cursor{ID: entity.NewID("1")},
					EndCursor:       &query.Cursor{ID: entity.NewID("1")},
				},
				Items: []*query.Recipe{
					{ID: entity.NewID("1")},
				},
			},
			result: &model.RecipeConnection{
				PageInfo: &model.PageInfo{
					HasNextPage:     true,
					HasPreviousPage: true,
					StartCursor:     &model.Cursor{ID: entity.NewID("1"), Sort: model.SortCreated},
					EndCursor:       &model.Cursor{ID: entity.NewID("1"), Sort: model.SortCreated},
				},
				Edges: []*model.RecipeEdge{
					{
						Cursor: model.Cursor{ID: entity.NewID("1"), Sort: model.SortCreated},
						Node:   model.NewRecipe(&query.Recipe{ID: entity.NewID("1")}),
					},
				},
			},
		},
		"updated sort": {
			page: &query.RecipePage{
				Info: query.PageInfo{
					HasNextPage:     true,
					HasPreviousPage: true,
					StartCursor:     &query.Cursor{ID: entity.NewID("1"), Sort: query.UpdatedSort},
				},
				Items: []*query.Recipe{
					{ID: entity.NewID("1")},
				},
			},
			result: &model.RecipeConnection{
				PageInfo: &model.PageInfo{
					HasNextPage:     true,
					HasPreviousPage: true,
					StartCursor:     &model.Cursor{ID: entity.NewID("1"), Sort: model.SortUpdated},
				},
				Edges: []*model.RecipeEdge{
					{
						Cursor: model.Cursor{ID: entity.NewID("1"), Sort: model.SortUpdated},
						Node:   model.NewRecipe(&query.Recipe{ID: entity.NewID("1")}),
					},
				},
			},
		},
		"name sort": {
			page: &query.RecipePage{
				Info: query.PageInfo{
					HasNextPage:     true,
					HasPreviousPage: true,
					StartCursor:     &query.Cursor{ID: entity.NewID("1"), Sort: query.NameSort},
					EndCursor:       &query.Cursor{ID: entity.NewID("3"), Sort: query.NameSort},
				},
				Items: []*query.Recipe{
					{ID: entity.NewID("1")},
					{ID: entity.NewID("2")},
					{ID: entity.NewID("3")},
				},
			},
			result: &model.RecipeConnection{
				PageInfo: &model.PageInfo{
					HasNextPage:     true,
					HasPreviousPage: true,
					StartCursor:     &model.Cursor{ID: entity.NewID("1"), Sort: model.SortName},
					EndCursor:       &model.Cursor{ID: entity.NewID("3"), Sort: model.SortName},
				},
				Edges: []*model.RecipeEdge{
					{
						Cursor: model.Cursor{ID: entity.NewID("1"), Sort: model.SortName},
						Node:   model.NewRecipe(&query.Recipe{ID: entity.NewID("1")}),
					},
					{
						Cursor: model.Cursor{ID: entity.NewID("2"), Sort: model.SortName},
						Node:   model.NewRecipe(&query.Recipe{ID: entity.NewID("2")}),
					},
					{
						Cursor: model.Cursor{ID: entity.NewID("3"), Sort: model.SortName},
						Node:   model.NewRecipe(&query.Recipe{ID: entity.NewID("3")}),
					},
				},
			},
		},
		"empty": {
			page: &query.RecipePage{},
			result: &model.RecipeConnection{
				PageInfo: &model.PageInfo{},
				Edges:    []*model.RecipeEdge{},
			},
		},
		"nil": {
			page: nil,
			result: &model.RecipeConnection{
				PageInfo: &model.PageInfo{},
				Edges:    []*model.RecipeEdge{},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.result, model.NewRecipeConnection(test.page))
		})
	}
}

func TestNewUserID(t *testing.T) {
	t.Parallel()

	id := entity.NewID("1234")

	result := model.ID{
		Key:  entity.NewID("1234"),
		Kind: model.UserKind,
	}

	assert.Equal(t, result, model.NewUserID(id))
}

func TestNewUser(t *testing.T) {
	t.Parallel()

	user := &query.User{
		ID:       entity.NewID("1234"),
		Username: "tester",
	}

	result := &model.User{
		ID:       model.NewUserID(entity.NewID("1234")),
		Username: "tester",
	}

	assert.Equal(t, result, model.NewUser(user))
}

func TestNewQueryPagination(t *testing.T) {
	t.Parallel()

	type testCase struct {
		page   *model.Page
		result query.Pagination
	}

	first := 100

	tests := map[string]testCase{
		"name sort": {
			page: &model.Page{
				First: &first,
				After: &model.Cursor{
					ID:   entity.NewID("1"),
					Sort: model.SortName,
				},
			},
			result: query.Pagination{
				Size: 100,
				Cursor: &query.Cursor{
					ID:   entity.NewID("1"),
					Sort: query.NameSort,
				},
			},
		},
		"created sort": {
			page: &model.Page{
				First: &first,
				After: &model.Cursor{
					ID:   entity.NewID("1"),
					Sort: model.SortCreated,
				},
			},
			result: query.Pagination{
				Size: 100,
				Cursor: &query.Cursor{
					ID:   entity.NewID("1"),
					Sort: query.CreatedSort,
				},
			},
		},
		"updated sort": {
			page: &model.Page{
				First: &first,
				After: &model.Cursor{
					ID:   entity.NewID("1"),
					Sort: model.SortUpdated,
				},
			},
			result: query.Pagination{
				Size: 100,
				Cursor: &query.Cursor{
					ID:   entity.NewID("1"),
					Sort: query.UpdatedSort,
				},
			},
		},
		"empty": {
			page: &model.Page{},
			result: query.Pagination{
				Size: 50,
			},
		},
		"nil": {
			page: nil,
			result: query.Pagination{
				Size: 50,
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.result, model.NewQueryPagination(test.page))
		})
	}
}

func TestNewQueryOrder(t *testing.T) {
	t.Parallel()

	type testCase struct {
		order  *model.Order
		result query.Order
	}

	sort := model.SortName
	dir := model.DirectionAsc

	tests := map[string]testCase{
		"value": {
			order: &model.Order{
				Sort:      &sort,
				Direction: &dir,
			},
			result: query.Order{
				Sort:      query.NameSort,
				Direction: query.AscDirection,
			},
		},
		"empty": {
			order:  &model.Order{},
			result: query.Order{},
		},
		"nil": {
			order:  nil,
			result: query.Order{},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.result, model.NewQueryOrder(test.order))
		})
	}
}
