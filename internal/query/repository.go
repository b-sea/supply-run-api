package query

import (
	"context"

	"github.com/b-sea/supply-run-api/internal/entity"
)

// Repository defines all data interactions required for querying recipes.
type Repository interface {
	FindRecipes(ctx context.Context, filter RecipeFilter, page Pagination, order Order) ([]*Recipe, error)
	GetRecipes(ctx context.Context, id []entity.ID) ([]*Recipe, error)
	GetIngredients(ctx context.Context) ([]string, error)
	GetTags(ctx context.Context) ([]string, error)
	GetUsers(ctx context.Context, ids []entity.ID) ([]*User, error)
}
