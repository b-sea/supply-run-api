package query

import "github.com/google/uuid"

type GetIngredientsHandler struct {
	reader GetIngredientsReader
}

func NewGetIngredientsHandler(reader GetIngredientsReader) *GetIngredientsHandler {
	return &GetIngredientsHandler{
		reader: reader,
	}
}

func (h GetIngredientsHandler) Handle(ingredientIDs []uuid.UUID) ([]*Ingredient, error) {
	return h.reader.GetIngredents(ingredientIDs)
}

type GetIngredientsReader interface {
	GetIngredents(ingredientIDs []uuid.UUID) ([]*Ingredient, error)
}
