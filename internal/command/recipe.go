package command

import (
	"context"

	"github.com/b-sea/supply-run-api/internal/entity"
	"github.com/b-sea/supply-run-api/internal/recipe"
	"github.com/bcicen/go-units"
)

type CreateRecipe struct {
	Name        string
	URL         string
	NumServings int
	Steps       []string
	Ingredients []Ingredient
	Tags        []string
}

type Ingredient struct {
	Name     string
	Quantity float64
	Unit     units.Unit
}

func (s *Service) CreateRecipe(ctx context.Context, cmd CreateRecipe) (entity.ID, error) {
	// TODO: Pull creation user id from context
	userID := s.idFn()

	options := []recipe.Option{
		recipe.SetURL(cmd.URL),
		recipe.SetNumServings(cmd.NumServings),
	}

	for _, step := range cmd.Steps {
		options = append(options, recipe.AddStep(step))
	}

	for _, ingredient := range cmd.Ingredients {
		options = append(options, recipe.AddIngredient(ingredient.Name, ingredient.Quantity, ingredient.Unit))
	}

	for _, tag := range cmd.Tags {
		options = append(options, recipe.AddTag(tag))
	}

	created, err := recipe.New(s.idFn(), cmd.Name, s.timestampFn(), userID, options...)
	if err != nil {
		return entity.ID{}, err
	}

	if err := s.recipes.CreateRecipe(ctx, created); err != nil {
		return entity.ID{}, err
	}

	return created.ID(), nil
}

type UpdateRecipe struct {
	ID          entity.ID
	Name        *string
	URL         *string
	NumServings *int
	Steps       []string
	Ingredients []Ingredient
	Tags        []string
}

func (s *Service) UpdateRecipe(ctx context.Context, cmd UpdateRecipe) error {
	// TODO: Pull creation user id from context
	userID := s.idFn()

	found, err := s.recipes.GetRecipe(ctx, cmd.ID)
	if err != nil {
		return err
	}

	// TODO: can users only edit things they create?

	options := make([]recipe.Option, 0)

	if cmd.Name != nil {
		options = append(options, recipe.SetName(*cmd.Name))
	}

	if cmd.URL != nil {
		options = append(options, recipe.SetURL(*cmd.URL))
	}

	if cmd.NumServings != nil {
		options = append(options, recipe.SetNumServings(*cmd.NumServings))
	}

	if cmd.Steps != nil {
		options = append(options, recipe.ClearSteps())

		for _, step := range cmd.Steps {
			options = append(options, recipe.AddStep(step))
		}
	}

	if cmd.Ingredients != nil {
		options = append(options, recipe.ClearIngredients())

		for _, ingredient := range cmd.Ingredients {
			options = append(options, recipe.AddIngredient(ingredient.Name, ingredient.Quantity, ingredient.Unit))
		}
	}

	if cmd.Tags != nil {
		options = append(options, recipe.ClearTags())

		for _, tag := range cmd.Tags {
			options = append(options, recipe.AddTag(tag))
		}
	}

	if err := found.Update(s.timestampFn(), userID, options...); err != nil {
		return err
	}

	if err := s.recipes.UpdateRecipe(ctx, found); err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteRecipe(ctx context.Context, id entity.ID) error {
	if err := s.recipes.DeleteRecipe(ctx, id); err != nil {
		return err
	}

	return nil
}
