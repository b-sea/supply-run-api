package query

import (
	"context"

	"github.com/b-sea/supply-run-api/internal/entity"
)

// RecipeRepository defines all data interactions required for querying recipes.
type RecipeRepository interface {
	FindRecipes(ctx context.Context, filter RecipeFilter, page Pagination, order Order) ([]*Recipe, error)
	GetRecipes(ctx context.Context, ids []entity.ID) ([]*Recipe, error)
	FindTags(ctx context.Context, filter *string) ([]string, error)
}

// UnitRepository defines all data interactions required for querying units.
type UnitRepository interface {
	GetUnits(ctx context.Context, ids []entity.ID) ([]*Unit, error)
}

// UserRepository defines all data interactions required for querying users.
type UserRepository interface {
	GetUsers(ctx context.Context, ids []entity.ID) ([]*User, error)
}
