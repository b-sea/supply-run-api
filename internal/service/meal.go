// Package service implements all business logic for the API.
package service

import (
	"github.com/b-sea/supply-run-api/internal/model"
	"github.com/b-sea/supply-run-api/internal/repository"
)

type IRecipe interface {
	INode[model.Recipe, model.RecipeFilter, model.CreateRecipeInput, model.UpdateRecipeInput]
}

type Recipe struct {
	node[model.Recipe, model.RecipeFilter, model.CreateRecipeInput, model.UpdateRecipeInput]
}

func NewRecipe(repo repository.IRecipe) *Recipe {
	return &Recipe{
		node: node[model.Recipe, model.RecipeFilter, model.CreateRecipeInput, model.UpdateRecipeInput]{
			repo:        repo,
			idGenerator: idGenerator,
			timestamp:   timestamp,
		},
	}
}
