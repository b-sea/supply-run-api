// Package recipe defines everything to manage the recipes domain.
package recipe

import (
	"time"

	"github.com/b-sea/supply-run-api/internal/entity"
	"github.com/google/uuid"
)

// Option is a recipe creation option.
type Option func(*Recipe)

// WithRecipeID sets the recipe id.
func WithRecipeID(id uuid.UUID) Option {
	return func(r *Recipe) {
		r.id = id
	}
}

// WithTimestamp sets the recipe creation time.
func WithTimestamp(now time.Time) Option {
	return func(r *Recipe) {
		r.createdAt = now
	}
}

// WithDescription sets the recipe description.
func WithDescription(desc string) Option {
	return func(r *Recipe) {
		r.Description = desc
	}
}

// WithURL sets the recipe URL.
func WithURL(url string) Option {
	return func(r *Recipe) {
		r.URL = url
	}
}

// WithServings sets the recipe description.
func WithServings(servings int) Option {
	return func(r *Recipe) {
		r.Servings = servings
	}
}

// WithIngredients sets the recipe ingredients.
func WithIngredients(ingredients []*Ingredient) Option {
	return func(r *Recipe) {
		r.Ingredients = ingredients
	}
}

// WithSteps sets the recipe steps.
func WithSteps(steps []*Step) Option {
	return func(r *Recipe) {
		r.Steps = steps
	}
}

// WithTags sets the recipe tags.
func WithTags(tags []*Tag) Option {
	return func(r *Recipe) {
		r.Tags = tags
	}
}

// Recipe is a cooking recipe.
type Recipe struct {
	id          uuid.UUID
	createdAt   time.Time
	owner       uuid.UUID
	Name        string
	Description string
	URL         string
	Servings    int
	Ingredients []*Ingredient
	Steps       []*Step
	Tags        []*Tag
}

// ID returns the recipe id.
func (r *Recipe) ID() uuid.UUID {
	return r.id
}

// CreatedAt returns a timestamp when the recipe is created.
func (r *Recipe) CreatedAt() time.Time {
	return r.createdAt
}

// Owner returns the creator of the unit.
func (r *Recipe) Owner() uuid.UUID {
	return r.owner
}

// Validate the recipe.
func (r *Recipe) Validate() error {
	issues := []string{}

	if r.Name == "" {
		issues = append(issues, "name cannot be empty")
	}

	if len(issues) == 0 {
		return nil
	}

	return &entity.ValidationError{
		Issues: issues,
	}
}

// NewRecipe creates a new recipe.
func NewRecipe(name string, owner uuid.UUID, opts ...Option) *Recipe {
	recipe := &Recipe{
		id:        uuid.New(),
		createdAt: time.Now().UTC(),
		owner:     owner,
		Name:      name,
		Servings:  1,
	}

	for _, opt := range opts {
		opt(recipe)
	}

	return recipe
}
