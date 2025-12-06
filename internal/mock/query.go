package mock

import (
	"context"

	"github.com/b-sea/supply-run-api/internal/entity"
	"github.com/b-sea/supply-run-api/internal/query"
)

var _ query.RecipeRepository = (*QueryRecipeRepository)(nil)
var _ query.UnitRepository = (*QueryUnitRepository)(nil)
var _ query.UserRepository = (*QueryUserRepository)(nil)

type QueryRecipeRepository struct {
	FindRecipesResult []*query.Recipe
	FindRecipesErr    error
	GetRecipesResult  []*query.Recipe
	GetRecipesErr     error
	FindTagsResult    []string
	FindTagsErr       error
}

func (m *QueryRecipeRepository) FindRecipes(
	ctx context.Context,
	filter query.RecipeFilter,
	page query.Pagination,
	order query.Order,
) ([]*query.Recipe, error) {
	return m.FindRecipesResult, m.FindRecipesErr
}

func (m *QueryRecipeRepository) GetRecipes(ctx context.Context, id []entity.ID) ([]*query.Recipe, error) {
	return m.GetRecipesResult, m.GetRecipesErr
}

func (m *QueryRecipeRepository) FindTags(ctx context.Context, filter *string) ([]string, error) {
	return m.FindTagsResult, m.FindTagsErr
}

type QueryUnitRepository struct {
	GetUnitsResult          []*query.Unit
	GetUnitsErr             error
	GetConversionPathResult []*query.Conversion
	GetConversionPathErr    error
	AllUnitsResult          []*query.Unit
	AllUnitsErr             error
}

func (m *QueryUnitRepository) GetUnits(ctx context.Context, ids []entity.ID) ([]*query.Unit, error) {
	return m.GetUnitsResult, m.GetUnitsErr
}

func (m *QueryUnitRepository) GetConversionPath(ctx context.Context, from *query.Unit, to *query.Unit) ([]*query.Conversion, error) {
	return m.GetConversionPathResult, m.GetConversionPathErr
}

func (m *QueryUnitRepository) AllUnits(ctx context.Context) ([]*query.Unit, error) {
	return m.AllUnitsResult, m.AllUnitsErr
}

type QueryUserRepository struct {
	GetUsersResult []*query.User
	GetUsersErr    error
}

func (m *QueryUserRepository) GetUsers(ctx context.Context, ids []entity.ID) ([]*query.User, error) {
	return m.GetUsersResult, m.GetUsersErr
}
