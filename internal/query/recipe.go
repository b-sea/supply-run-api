package query

import (
	"context"
	"slices"

	"github.com/b-sea/supply-run-api/internal/entity"
)

// FindRecipes returns a list of recipes based on search criteria.
func (s *Service) FindRecipes(ctx context.Context, filter *RecipeFilter, page *Pagination) ([]*Recipe, error) {
	if filter == nil {
		filter = &RecipeFilter{}
	}

	if page == nil {
		page = &Pagination{}
	}

	result, err := s.repo.FindRecipes(ctx, *filter, *page)
	if err != nil {
		return nil, queryError(err)
	}

	return result, nil
}

// GetRecipe returns a single recipe from an id.
func (s *Service) GetRecipe(ctx context.Context, id entity.ID) (*Recipe, error) {
	found, err := s.repo.GetRecipes(ctx, []entity.ID{id})
	if err != nil {
		return nil, queryError(err)
	}

	if len(found) == 0 {
		return nil, entity.ErrNotFound
	}

	return found[0], nil
}

// GetIngredients returns a case-insensitive list of unique ingredients found on recipes.
func (s *Service) GetIngredients(ctx context.Context) ([]string, error) {
	found, err := s.repo.GetIngredients(ctx)
	if err != nil {
		return nil, queryError(err)
	}

	slices.Sort(found)

	return found, nil
}

// GetTags returns a case-insensitive list of unique tags found on recipes.
func (s *Service) GetTags(ctx context.Context) ([]string, error) {
	found, err := s.repo.GetTags(ctx)
	if err != nil {
		return nil, queryError(err)
	}

	slices.Sort(found)

	return found, nil
}
