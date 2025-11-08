package recipe

import "github.com/b-sea/supply-run-api/internal/entity"

// Repository defines all data interactions required for recipes.
type Repository interface {
	GetRecipes(id []entity.ID) ([]*Recipe, error)
	CreateRecipe(recipe *Recipe) error
	UpdateRecipe(recipe *Recipe) error
	DeleteRecipe(id entity.ID) error
}
