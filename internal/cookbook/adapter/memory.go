package adapter

import (
	"slices"
	"strings"

	"github.com/b-sea/supply-run-api/internal/cookbook/app/query"
	"github.com/b-sea/supply-run-api/internal/cookbook/domain/recipe"
	"github.com/b-sea/supply-run-api/internal/cookbook/domain/unit"
	"github.com/google/uuid"
)

type CookbookMemoryRepository struct {
	recipes []*recipe.Recipe
	units   []*unit.Unit
}

func NewCookbookMemoryRepository() *CookbookMemoryRepository {
	return &CookbookMemoryRepository{
		recipes: make([]*recipe.Recipe, 0),
		units:   make([]*unit.Unit, 0),
	}
}

func (r *CookbookMemoryRepository) FindRecipeSnippets(
	userID uuid.UUID,
	filter *query.RecipeFilter,
) ([]*query.RecipeSnippet, error) {
	results := make([]*query.RecipeSnippet, 0)

	for _, rcp := range r.recipes {
		if rcp.OwnerID() != userID {
			continue
		}

		if r.recipeIsFiltered(rcp, filter) {
			continue
		}

		results = append(results, r.recipeToRecipeSnippet(rcp))
	}

	return results, nil
}

func (r *CookbookMemoryRepository) recipeIsFiltered(rcp *recipe.Recipe, filter *query.RecipeFilter) bool {
	if filter == nil {
		return false
	}

	if filter.Name != nil {
		return strings.Contains(strings.ToLower(rcp.Name()), strings.ToLower(*filter.Name))
	}

	// TODO: handle this filter
	if len(filter.TagIDs) > 0 {
	}

	return false
}

func (r *CookbookMemoryRepository) recipeToRecipeSnippet(rcp *recipe.Recipe) *query.RecipeSnippet {
	return &query.RecipeSnippet{
		ID:          rcp.ID(),
		CreatedAt:   rcp.CreatedAt(),
		Name:        rcp.Name(),
		Description: rcp.Description(),
		TagIDs:      rcp.TagIDs(),
	}
}

func (r *CookbookMemoryRepository) GetRecipe(userID uuid.UUID, recipeID uuid.UUID) (*query.Recipe, error) {
	for _, rcp := range r.recipes {
		if rcp.OwnerID() != userID {
			continue
		}

		if rcp.ID() != recipeID {
			continue
		}

		return r.recipeToQueryRecipe(rcp), nil
	}

	return nil, ErrNotFound
}

func (r *CookbookMemoryRepository) recipeToQueryRecipe(rcp *recipe.Recipe) *query.Recipe {
	result := &query.Recipe{
		ID:            rcp.ID(),
		CreatedAt:     rcp.CreatedAt(),
		UpdatedAt:     rcp.UpdatedAt(),
		OwnerID:       rcp.OwnerID(),
		Name:          rcp.Name(),
		Description:   rcp.Description(),
		URL:           rcp.URL(),
		IngredientIDs: make([]uuid.UUID, 0),
		Servings:      rcp.Servings(),
		Steps:         rcp.Steps(),
		TagIDs:        rcp.TagIDs(),
	}

	for _, ingredient := range rcp.Ingredients() {
		result.IngredientIDs = append(result.IngredientIDs, ingredient.ID())
	}

	return result
}

func (r *CookbookMemoryRepository) GetIngredents(ingredientIDs []uuid.UUID) ([]*query.Ingredient, error) {
	result := make([]*query.Ingredient, 0)

	for _, rcp := range r.recipes {
		for _, ingredient := range rcp.Ingredients() {
			if !slices.Contains(ingredientIDs, ingredient.ID()) {
				continue
			}

			result = append(result, r.ingredientToQueryIngredient(ingredient))
		}
	}

	return result, nil
}

func (r *CookbookMemoryRepository) ingredientToQueryIngredient(ingredient *recipe.Ingredient) *query.Ingredient {
	return &query.Ingredient{
		ID:       ingredient.ID(),
		ItemID:   ingredient.ItemID(),
		UnitID:   ingredient.UnitID(),
		Quantity: ingredient.Quantity(),
	}
}

func (r *CookbookMemoryRepository) GetAllUnits(userID uuid.UUID) ([]*query.Unit, error) {
	result := make([]*query.Unit, 0)

	for _, measurement := range r.units {
		if userID != measurement.OwnerID() {
			continue
		}

		result = append(result, r.unitToQueryUnit(measurement))
	}

	return result, nil
}

func (r *CookbookMemoryRepository) GetUnits(userID uuid.UUID, unitIDs []uuid.UUID) ([]*query.Unit, error) {
	result := make([]*query.Unit, 0)

	for _, measurement := range r.units {
		if userID != measurement.OwnerID() {
			continue
		}

		if !slices.Contains(unitIDs, measurement.ID()) {
			continue
		}

		result = append(result, r.unitToQueryUnit(measurement))
	}

	return result, nil
}

func (r *CookbookMemoryRepository) unitToQueryUnit(measurement *unit.Unit) *query.Unit {
	return &query.Unit{
		ID:        measurement.ID(),
		CreatedAt: measurement.CreatedAt(),
		UpdatedAt: measurement.UpdatedAt(),
		OwnerID:   measurement.OwnerID(),
		Name:      measurement.Name(),
		Symbol:    measurement.Symbol(),
		Type:      measurement.Type(),
		System:    measurement.System(),
	}
}
