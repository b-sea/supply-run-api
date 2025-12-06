package mariadb

import (
	"context"

	"github.com/b-sea/supply-run-api/internal/entity"
	"github.com/b-sea/supply-run-api/internal/query"
)

func (r *Repository) FindRecipes(
	ctx context.Context,
	filter query.RecipeFilter,
	page query.Pagination,
	order query.Order,
) ([]*query.Recipe, error) {
	return nil, nil
}

func (r *Repository) GetRecipes(ctx context.Context, ids []entity.ID) ([]*query.Recipe, error) {
	return nil, nil
}

func (r *Repository) FindTags(ctx context.Context, filter *string) ([]string, error) {
	return nil, nil
}
