package query

import (
	"context"
	"slices"

	"github.com/b-sea/supply-run-api/internal/entity"
)

const pagePadding = 2

// FindRecipes returns a list of recipes based on search criteria.
func (s *Service) FindRecipes(
	ctx context.Context,
	filter RecipeFilter,
	page Pagination,
	order Order,
) (*RecipePage, error) {
	result := &RecipePage{
		Info:  PageInfo{},
		Items: make([]*Recipe, 0),
	}

	if page.Size == 0 {
		return result, nil
	}

	// Use the +2 trick: https://stackoverflow.com/a/66300422
	pageSize := page.Size
	page.Size += pagePadding

	found, err := s.repo.FindRecipes(ctx, filter, page, order)
	if err != nil {
		return nil, queryError(err)
	}

	if len(found) == 0 {
		return result, nil
	}

	sort := CreatedSort

	if page.Cursor != nil && page.Cursor.ID == found[0].ID {
		sort = page.Cursor.Sort
		result.Info.HasPreviousPage = true
		found = found[1:]
	}

	count := len(found)
	result.Info.HasNextPage = count > pageSize

	if count == 0 {
		return result, nil
	}

	if pageSize > count {
		pageSize = count
	}

	for i := range pageSize {
		result.Items = append(result.Items, found[i])
	}

	result.Info.StartCursor = &Cursor{
		ID:   result.Items[0].ID,
		Sort: sort,
	}

	result.Info.EndCursor = &Cursor{
		ID:   result.Items[len(result.Items)-1].ID,
		Sort: sort,
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
