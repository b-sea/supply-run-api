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

func TestQueryNode(t *testing.T) {
	t.Parallel()

	type testCase struct {
		recipe   query.RecipeRepository
		unit     query.UnitRepository
		user     query.UserRepository
		options  []client.Option
		query    string
		response map[string]any
		err      error
	}

	tests := map[string]testCase{
		"recipe found": {
			recipe: &mock.QueryRecipeRepository{
				GetRecipesResult: []*query.Recipe{
					{ID: entity.NewID("1")},
				},
			},
			unit:    &mock.QueryUnitRepository{},
			user:    &mock.QueryUserRepository{},
			options: []client.Option{client.Var("id", model.NewRecipeID(entity.NewID("1")).String())},
			query:   `query test($id: ID!){ node(id: $id){ __typename }}`,
			response: map[string]any{
				"node": map[string]any{
					"__typename": "Recipe",
				},
			},
			err: nil,
		},
		"recipe not found": {
			recipe:  &mock.QueryRecipeRepository{},
			unit:    &mock.QueryUnitRepository{},
			user:    &mock.QueryUserRepository{},
			options: []client.Option{client.Var("id", model.NewRecipeID(entity.NewID("1")).String())},
			query:   `query test($id: ID!){ node(id: $id){ __typename }}`,
			response: map[string]any{
				"node": map[string]any{
					"__typename": "NotFoundError",
				},
			},
			err: nil,
		},
		"recipe error": {
			recipe: &mock.QueryRecipeRepository{
				GetRecipesErr: errors.New("some random error"),
			},
			unit:     &mock.QueryUnitRepository{},
			user:     &mock.QueryUserRepository{},
			options:  []client.Option{client.Var("id", model.NewRecipeID(entity.NewID("1")).String())},
			query:    `query test($id: ID!){ node(id: $id){ __typename }}`,
			response: nil,
			err:      errors.New("some random error"),
		},
		"unit found": {
			recipe: &mock.QueryRecipeRepository{},
			unit: &mock.QueryUnitRepository{
				GetUnitsResult: []*query.Unit{
					{ID: entity.NewID("1")},
				},
			},
			user:    &mock.QueryUserRepository{},
			options: []client.Option{client.Var("id", model.NewUnitID(entity.NewID("1")).String())},
			query:   `query test($id: ID!){ node(id: $id){ __typename }}`,
			response: map[string]any{
				"node": map[string]any{
					"__typename": "Unit",
				},
			},
			err: nil,
		},
		"unit not found": {
			recipe:  &mock.QueryRecipeRepository{},
			unit:    &mock.QueryUnitRepository{},
			user:    &mock.QueryUserRepository{},
			options: []client.Option{client.Var("id", model.NewUnitID(entity.NewID("1")).String())},
			query:   `query test($id: ID!){ node(id: $id){ __typename }}`,
			response: map[string]any{
				"node": map[string]any{
					"__typename": "NotFoundError",
				},
			},
			err: nil,
		},
		"unit error": {
			recipe: &mock.QueryRecipeRepository{},
			unit: &mock.QueryUnitRepository{
				GetUnitsErr: errors.New("some random error"),
			},
			user:     &mock.QueryUserRepository{},
			options:  []client.Option{client.Var("id", model.NewUnitID(entity.NewID("1")).String())},
			query:    `query test($id: ID!){ node(id: $id){ __typename }}`,
			response: nil,
			err:      errors.New("some random error"),
		},
		"user found": {
			recipe: &mock.QueryRecipeRepository{},
			unit:   &mock.QueryUnitRepository{},
			user: &mock.QueryUserRepository{
				GetUsersResult: []*query.User{
					{ID: entity.NewID("1")},
				},
			},
			options: []client.Option{client.Var("id", model.NewUserID(entity.NewID("1")).String())},
			query:   `query test($id: ID!){ node(id: $id){ __typename }}`,
			response: map[string]any{
				"node": map[string]any{
					"__typename": "User",
				},
			},
			err: nil,
		},
		"user not found": {
			recipe:  &mock.QueryRecipeRepository{},
			unit:    &mock.QueryUnitRepository{},
			user:    &mock.QueryUserRepository{},
			options: []client.Option{client.Var("id", model.NewUserID(entity.NewID("1")).String())},
			query:   `query test($id: ID!){ node(id: $id){ __typename }}`,
			response: map[string]any{
				"node": map[string]any{
					"__typename": "NotFoundError",
				},
			},
			err: nil,
		},
		"user error": {
			recipe: &mock.QueryRecipeRepository{},
			unit:   &mock.QueryUnitRepository{},
			user: &mock.QueryUserRepository{
				GetUsersErr: errors.New("some random error"),
			},
			options:  []client.Option{client.Var("id", model.NewUserID(entity.NewID("1")).String())},
			query:    `query test($id: ID!){ node(id: $id){ __typename }}`,
			response: nil,
			err:      errors.New("some random error"),
		},
		"unknown type": {
			recipe:  &mock.QueryRecipeRepository{},
			unit:    &mock.QueryUnitRepository{},
			user:    &mock.QueryUserRepository{},
			options: []client.Option{client.Var("id", model.ID{Key: entity.NewID("1"), Kind: model.Kind("random")}.String())},
			query:   `query test($id: ID!){ node(id: $id){ __typename }}`,
			response: map[string]any{
				"node": map[string]any{
					"__typename": "NotFoundError",
				},
			},
			err: nil,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			server := graphql.New(
				query.NewService(test.recipe, test.unit, test.user),
				metrics.NewNoOp(),
			)
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
