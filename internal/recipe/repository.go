package recipe

import (
	"context"

	"github.com/b-sea/supply-run-api/internal/entity"
)

// Repository defines all data interactions required for recipes.
type Repository interface {
	GetRecipe(ctx context.Context, id entity.ID) (*Recipe, error)
	CreateRecipe(ctx context.Context, recipe *Recipe) error
	UpdateRecipe(ctx context.Context, recipe *Recipe) error
	DeleteRecipe(ctx context.Context, id entity.ID) error
}
