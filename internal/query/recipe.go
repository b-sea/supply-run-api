package query

import (
	"context"
	"slices"

	"github.com/b-sea/go-logger/logger"
	"github.com/b-sea/supply-run-api/internal/entity"
)

// GetRecipe returns a single recipe from an id.
func (s *Service) GetRecipe(ctx context.Context, id entity.ID) (*Recipe, error) {
	logger.FromContext(ctx).Trace().Str("id", id.String()).Msg("get recipe")

	found, err := s.GetRecipes(ctx, []entity.ID{id})
	if err != nil {
		return nil, err
	}

	if len(found) == 0 {
		return nil, entity.ErrNotFound
	}

	return found[0], nil
}

// GetRecipes returns multiple recipes from a list of ids.
func (s *Service) GetRecipes(ctx context.Context, ids []entity.ID) ([]*Recipe, error) {
	logger.FromContext(ctx).Trace().Int("count", len(ids)).Msg("get recipes")

	found, err := s.repo.GetRecipes(ctx, ids)
	if err != nil {
		return nil, queryError(err)
	}

	return found, nil
}

// GetIngredients returns a case-insensitive list of unique ingredients found on recipes.
func (s *Service) GetIngredients(ctx context.Context) ([]string, error) {
	logger.FromContext(ctx).Trace().Msg("get ingredients")

	found, err := s.repo.GetIngredients(ctx)
	if err != nil {
		return nil, queryError(err)
	}

	slices.Sort(found)

	return found, nil
}

// GetTags returns a case-insensitive list of unique tags found on recipes.
func (s *Service) GetTags(ctx context.Context) ([]string, error) {
	logger.FromContext(ctx).Trace().Msg("get tags")

	found, err := s.repo.GetTags(ctx)
	if err != nil {
		return nil, queryError(err)
	}

	slices.Sort(found)

	return found, nil
}
