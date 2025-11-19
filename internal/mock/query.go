package mock

import (
	"context"

	"github.com/b-sea/supply-run-api/internal/entity"
	"github.com/b-sea/supply-run-api/internal/query"
)

var _ query.Repository = (*QueryRepository)(nil)

type QueryRepository struct {
	FindRecipesResult   []*query.Recipe
	FindRecipesErr      error
	GetRecipesResult    []*query.Recipe
	GetRecipesErr       error
	AllRecipeTagsResult []string
	AllRecipeTagsErr    error
	GetUsersResult      []*query.User
	GetUsersErr         error
}

func (m *QueryRepository) FindRecipes(
	ctx context.Context,
	filter query.RecipeFilter,
	page query.Pagination,
	order query.Order,
) ([]*query.Recipe, error) {
	return m.FindRecipesResult, m.FindRecipesErr
}

func (m *QueryRepository) GetRecipes(ctx context.Context, id []entity.ID) ([]*query.Recipe, error) {
	return m.GetRecipesResult, m.GetRecipesErr
}

func (m *QueryRepository) AllRecipeTags(ctx context.Context) ([]string, error) {
	return m.AllRecipeTagsResult, m.AllRecipeTagsErr
}

func (m *QueryRepository) GetUsers(ctx context.Context, ids []entity.ID) ([]*query.User, error) {
	return m.GetUsersResult, m.GetUsersErr
}
