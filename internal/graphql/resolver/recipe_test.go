package resolver_test

import (
	"errors"
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/b-sea/supply-run-api/internal/entity"
	"github.com/b-sea/supply-run-api/internal/graphql"
	"github.com/b-sea/supply-run-api/internal/graphql/model"
	"github.com/b-sea/supply-run-api/internal/metrics"
	"github.com/b-sea/supply-run-api/internal/mock"
	"github.com/b-sea/supply-run-api/internal/query"
	"github.com/stretchr/testify/assert"
)

func TestQueryFindRecipes(t *testing.T) {
	t.Parallel()

	type testCase struct {
		repo     query.Repository
		options  []client.Option
		query    string
		response map[string]any
		err      error
	}

	tests := map[string]testCase{
		"success": {
			repo: &mock.QueryRepository{
				FindRecipesResult: []*query.Recipe{
					{ID: entity.NewID("1")},
				},
			},
			query: `query { findRecipes { pageInfo { hasNextPage hasPreviousPage startCursor endCursor } edges { cursor node { id }}}}`,
			response: map[string]any{
				"findRecipes": map[string]any{
					"edges": []any{
						map[string]any{"cursor": "MTpDUkVBVEVE", "node": map[string]any{"id": "cmVjaXBlOjE="}},
					},
					"pageInfo": map[string]any{
						"endCursor":       "MTpDUkVBVEVE",
						"hasNextPage":     false,
						"hasPreviousPage": false,
						"startCursor":     "MTpDUkVBVEVE",
					},
				},
			},
			err: nil,
		},
		"repo error": {
			repo: &mock.QueryRepository{
				FindRecipesErr: errors.New("some random error"),
			},
			query:    `query { findRecipes { pageInfo { hasNextPage hasPreviousPage startCursor endCursor } edges { cursor node { id }}}}`,
			response: nil,
			err:      errors.New("some random error"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			server := graphql.New(query.NewService(test.repo), metrics.NewNoOp())
			testClient := client.New(server)

			var response map[string]any

			err := testClient.Post(test.query, &response, test.options...)

			assert.Equal(t, test.response, response)
			if test.err == nil {
				assert.NoError(t, err)
			} else {
				assert.ErrorAs(t, err, &test.err)
			}
		})
	}
}

func TestQueryRecipe(t *testing.T) {
	t.Parallel()

	type testCase struct {
		repo     query.Repository
		options  []client.Option
		query    string
		response map[string]any
		err      error
	}

	tests := map[string]testCase{
		"success": {
			repo: &mock.QueryRepository{
				GetRecipesResult: []*query.Recipe{
					{ID: entity.NewID("1")},
				},
			},
			options: []client.Option{client.Var("id", model.NewRecipeID(entity.NewID("1")).String())},
			query:   `query recipeByID($id: ID!){ recipe(id: $id) { __typename ...on Recipe { id }}}`,
			response: map[string]any{
				"recipe": map[string]any{
					"__typename": "Recipe",
					"id":         model.NewRecipeID(entity.NewID("1")).String(),
				},
			},
			err: nil,
		},
		"not found": {
			repo: &mock.QueryRepository{
				GetRecipesResult: []*query.Recipe{},
			},
			options: []client.Option{client.Var("id", model.NewRecipeID(entity.NewID("1")).String())},
			query:   `query recipeByID($id: ID!){ recipe(id: $id) { __typename ...on NotFoundError { id }}}`,
			response: map[string]any{
				"recipe": map[string]any{
					"__typename": "NotFoundError",
					"id":         model.NewRecipeID(entity.NewID("1")).String(),
				},
			},
			err: nil,
		},
		"bad id": {
			repo: &mock.QueryRepository{
				GetRecipesResult: []*query.Recipe{},
			},
			options: []client.Option{client.Var("id", model.NewUserID(entity.NewID("1")).String())},
			query:   `query recipeByID($id: ID!){ recipe(id: $id) { __typename ...on NotFoundError { id }}}`,
			response: map[string]any{
				"recipe": map[string]any{
					"__typename": "NotFoundError",
					"id":         model.NewUserID(entity.NewID("1")).String(),
				},
			},
			err: nil,
		},
		"repo error": {
			repo: &mock.QueryRepository{
				GetRecipesErr: errors.New("some random error"),
			},
			options:  []client.Option{client.Var("id", model.NewRecipeID(entity.NewID("1")).String())},
			query:    `query recipeByID($id: ID!){ recipe(id: $id) { __typename ...on Recipe { id }}}`,
			response: nil,
			err:      errors.New("some random error"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			server := graphql.New(query.NewService(test.repo), metrics.NewNoOp())
			testClient := client.New(server)

			var response map[string]any

			err := testClient.Post(test.query, &response, test.options...)

			assert.Equal(t, test.response, response)
			if test.err == nil {
				assert.NoError(t, err)
			} else {
				assert.ErrorAs(t, err, &test.err)
			}
		})
	}
}

func TestQueryIngredients(t *testing.T) {
	t.Parallel()

	type testCase struct {
		repo     query.Repository
		options  []client.Option
		query    string
		response map[string]any
		err      error
	}

	tests := map[string]testCase{
		"success": {
			repo: &mock.QueryRepository{
				GetIngredientsResult: []string{
					"bread",
					"mustard",
					"cheese",
				},
			},
			query: `query { ingredients }`,
			response: map[string]any{
				"ingredients": []any{
					"bread",
					"cheese",
					"mustard",
				},
			},
			err: nil,
		},
		"repo error": {
			repo: &mock.QueryRepository{
				GetIngredientsErr: errors.New("some random error"),
			},
			query:    `query { ingredients }`,
			response: nil,
			err:      errors.New("some random error"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			server := graphql.New(query.NewService(test.repo), metrics.NewNoOp())
			testClient := client.New(server)

			var response map[string]any

			err := testClient.Post(test.query, &response, test.options...)

			assert.Equal(t, test.response, response)
			if test.err == nil {
				assert.NoError(t, err)
			} else {
				assert.ErrorAs(t, err, &test.err)
			}
		})
	}
}

func TestQueryTags(t *testing.T) {
	t.Parallel()

	type testCase struct {
		repo     query.Repository
		options  []client.Option
		query    string
		response map[string]any
		err      error
	}

	tests := map[string]testCase{
		"success": {
			repo: &mock.QueryRepository{
				GetTagsResult: []string{
					"vegan",
					"breakfast",
				},
			},
			query: `query { tags }`,
			response: map[string]any{
				"tags": []any{
					"breakfast",
					"vegan",
				},
			},
			err: nil,
		},
		"repo error": {
			repo: &mock.QueryRepository{
				GetTagsErr: errors.New("some random error"),
			},
			query:    `query { tags }`,
			response: nil,
			err:      errors.New("some random error"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			server := graphql.New(query.NewService(test.repo), metrics.NewNoOp())
			testClient := client.New(server)

			var response map[string]any

			err := testClient.Post(test.query, &response, test.options...)

			assert.Equal(t, test.response, response)
			if test.err == nil {
				assert.NoError(t, err)
			} else {
				assert.ErrorAs(t, err, &test.err)
			}
		})
	}
}

func TestQueryRecipeCreatedBy(t *testing.T) {
	t.Parallel()

	type testCase struct {
		repo     query.Repository
		options  []client.Option
		query    string
		response map[string]any
		err      error
	}

	tests := map[string]testCase{
		"success": {
			repo: &mock.QueryRepository{
				GetRecipesResult: []*query.Recipe{{ID: entity.NewID("R1"), CreatedBy: entity.NewID("U1")}},
				GetUsersResult:   []*query.User{{ID: entity.NewID("U1")}},
			},
			options: []client.Option{client.Var("id", model.NewRecipeID(entity.NewID("R1")).String())},
			query:   `query recipeByID($id: ID!){ recipe(id: $id) { ...on Recipe { createdBy { __typename ...on User { id }}}}}`,
			response: map[string]any{
				"recipe": map[string]any{
					"createdBy": map[string]any{
						"__typename": "User",
						"id":         "dXNlcjpVMQ==",
					},
				},
			},
			err: nil,
		},
		"repo error": {
			repo: &mock.QueryRepository{
				GetRecipesResult: []*query.Recipe{{ID: entity.NewID("R1"), CreatedBy: entity.NewID("U1")}},
				GetUsersErr:      errors.New("some random error"),
			},
			options:  []client.Option{client.Var("id", model.NewRecipeID(entity.NewID("1")).String())},
			query:    `query recipeByID($id: ID!){ recipe(id: $id) { ...on Recipe { createdBy { __typename ...on User { id }}}}}`,
			response: nil,
			err:      errors.New("some random error"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			server := graphql.New(query.NewService(test.repo), metrics.NewNoOp())
			testClient := client.New(server)

			var response map[string]any

			err := testClient.Post(test.query, &response, test.options...)

			assert.Equal(t, test.response, response)
			if test.err == nil {
				assert.NoError(t, err)
			} else {
				assert.ErrorAs(t, err, &test.err)
			}
		})
	}
}

func TestQueryRecipeUpdatedBy(t *testing.T) {
	t.Parallel()

	type testCase struct {
		repo     query.Repository
		options  []client.Option
		query    string
		response map[string]any
		err      error
	}

	tests := map[string]testCase{
		"success": {
			repo: &mock.QueryRepository{
				GetRecipesResult: []*query.Recipe{{ID: entity.NewID("R1"), UpdatedBy: entity.NewID("U1")}},
				GetUsersResult:   []*query.User{{ID: entity.NewID("U1")}},
			},
			options: []client.Option{client.Var("id", model.NewRecipeID(entity.NewID("R1")).String())},
			query:   `query recipeByID($id: ID!){ recipe(id: $id) { ...on Recipe { updatedBy { __typename ...on User { id }}}}}`,
			response: map[string]any{
				"recipe": map[string]any{
					"updatedBy": map[string]any{
						"__typename": "User",
						"id":         "dXNlcjpVMQ==",
					},
				},
			},
			err: nil,
		},
		"repo error": {
			repo: &mock.QueryRepository{
				GetRecipesResult: []*query.Recipe{{ID: entity.NewID("R1"), UpdatedBy: entity.NewID("U1")}},
				GetUsersErr:      errors.New("some random error"),
			},
			options:  []client.Option{client.Var("id", model.NewRecipeID(entity.NewID("1")).String())},
			query:    `query recipeByID($id: ID!){ recipe(id: $id) { ...on Recipe { updatedBy { __typename ...on User { id }}}}}`,
			response: nil,
			err:      errors.New("some random error"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			server := graphql.New(query.NewService(test.repo), metrics.NewNoOp())
			testClient := client.New(server)

			var response map[string]any

			err := testClient.Post(test.query, &response, test.options...)

			assert.Equal(t, test.response, response)
			if test.err == nil {
				assert.NoError(t, err)
			} else {
				assert.ErrorAs(t, err, &test.err)
			}
		})
	}
}
