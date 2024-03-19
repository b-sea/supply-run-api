// Package memory implements domains in memory storage.
package memory

import (
	"slices"
	"strings"

	"github.com/b-sea/supply-run-api/internal/domain/recipe"
	"github.com/b-sea/supply-run-api/internal/entity"
	"github.com/google/uuid"
)

// RecipeRepository implements the recipe domain.
type RecipeRepository struct {
	recipes map[uuid.UUID]*recipe.Recipe
}

// NewRecipeRepository creates a new RecipeRepository.
func NewRecipeRepository(recipes []*recipe.Recipe) *RecipeRepository {
	repo := RecipeRepository{
		recipes: make(map[uuid.UUID]*recipe.Recipe),
	}

	for _, r := range recipes {
		repo.recipes[r.ID()] = r
	}

	return &repo
}

func tagsIntersect(a []*recipe.Tag, b []*recipe.Tag) []*recipe.Tag { //nolint: varnamelen
	set := make([]*recipe.Tag, 0)
	hash := make(map[uuid.UUID]struct{})

	for _, v := range a {
		hash[v.ID()] = struct{}{}
	}

	for _, v := range b {
		if _, ok := hash[v.ID()]; ok {
			set = append(set, v)
		}
	}

	return set
}

func (r *RecipeRepository) isValid(entity *recipe.Recipe, filter *recipe.Filter) bool {
	if filter == nil {
		return true
	}

	if filter.Name != nil {
		if !strings.Contains(entity.Name, *filter.Name) {
			return false
		}
	}

	if len(filter.Tags) > 0 {
		if len(tagsIntersect(entity.Tags, filter.Tags)) == 0 {
			return false
		}
	}

	return true
}

// Find all recipes based on the given owners and filters.
func (r *RecipeRepository) Find(owners []uuid.UUID, filter *recipe.Filter) ([]*recipe.Recipe, error) {
	results := []*recipe.Recipe{}

	for _, entity := range r.recipes {
		if !slices.Contains(owners, entity.Owner()) {
			continue
		}

		if !r.isValid(entity, filter) {
			continue
		}

		results = append(results, entity)
	}

	return results, nil
}

// GetOne recipe.
func (r *RecipeRepository) GetOne(owners []uuid.UUID, id uuid.UUID) (*recipe.Recipe, error) {
	found := r.recipes[id]
	if found == nil {
		return nil, &entity.NotFoundError{
			ID: id,
		}
	}

	if !slices.Contains(owners, found.Owner()) {
		return nil, &entity.NotFoundError{
			ID: id,
		}
	}

	return found, nil
}

// Create a new recipe.
func (r *RecipeRepository) Create(entity *recipe.Recipe) error {
	r.recipes[entity.ID()] = entity
	return nil
}

// Update an existing recipe.
func (r *RecipeRepository) Update(entity *recipe.Recipe) error {
	r.recipes[entity.ID()] = entity
	return nil
}

// Delete an existing recipe.
func (r *RecipeRepository) Delete(id uuid.UUID) error {
	delete(r.recipes, id)
	return nil
}
