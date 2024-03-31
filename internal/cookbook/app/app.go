package app

import "github.com/b-sea/supply-run-api/internal/cookbook/app/query"

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
}

type Queries struct {
	FindRecipeSnippets query.FindRecipeSnippetsHandler
	GetRecipe          query.GetRecipeHandler
	GetIngredients     query.GetIngredientsHandler
	GetAllUnits        query.GetAllUnitsHandler
	GetUnits           query.GetUnitsHandler
}
