package query

import (
	"github.com/google/uuid"
)

type FindRecipeSnippetsHandler struct {
	reader FindRecipeSnippetsReader
}

func NewFindRecipeSnippetsHandler(reader FindRecipeSnippetsReader) *FindRecipeSnippetsHandler {
	return &FindRecipeSnippetsHandler{
		reader: reader,
	}
}

func (h *FindRecipeSnippetsHandler) Handle(userID uuid.UUID, filter *RecipeFilter) ([]*RecipeSnippet, error) {
	return h.reader.FindRecipeSnippets(userID, filter)
}

type FindRecipeSnippetsReader interface {
	FindRecipeSnippets(userID uuid.UUID, filter *RecipeFilter) ([]*RecipeSnippet, error)
}
