package query

import (
	"github.com/google/uuid"
)

type GetRecipeHandler struct {
	reader GetRecipeReader
}

func (h GetRecipeHandler) Handle(userID uuid.UUID, recipeID uuid.UUID) (*Recipe, error) {
	return h.reader.GetRecipe(userID, recipeID)
}

type GetRecipeReader interface {
	GetRecipe(userID uuid.UUID, recipeID uuid.UUID) (*Recipe, error)
}
