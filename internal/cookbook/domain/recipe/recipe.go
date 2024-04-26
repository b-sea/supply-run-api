// Package recipe defines everything to manage the recipes domain.
package recipe

import (
	"errors"
	"time"

	"github.com/b-sea/supply-run-api/internal/cookbook/domain"
	"github.com/google/uuid"
)

// Option is a recipe creation option.
type Option func(*Recipe) error

// SetName sets the recipe name.
func SetName(name string) Option {
	return func(r *Recipe) error {
		if name == "" {
			return errors.New("recipe name cannot be empty") //nolint: goerr113
		}

		r.name = name

		return nil
	}
}

// SetDescription sets the recipe description.
func SetDescription(desc string) Option {
	return func(r *Recipe) error {
		r.desc = desc
		return nil
	}
}

// SetURL sets the recipe URL.
func SetURL(url *string) Option {
	return func(r *Recipe) error {
		r.url = url
		return nil
	}
}

// SetServings sets the recipe description.
func SetServings(servings int) Option {
	return func(r *Recipe) error {
		if servings < 1 {
			return errors.New("recipe must have at least 1 serving") //nolint: goerr113
		}

		r.servings = servings

		return nil
	}
}

// SetSteps sets the recipe steps.
func SetSteps(steps []string) Option {
	return func(r *Recipe) error { //nolint: varnamelen
		for _, s := range steps {
			if s == "" {
				return errors.New("steps cannot be empty") //nolint: goerr113
			}
		}

		r.steps = steps

		return nil
	}
}

// SetIngredients sets the recipe ingredients.
func SetIngredients(ingredients []*Ingredient) Option {
	return func(r *Recipe) error {
		r.ingredients = ingredients
		return nil
	}
}

// SetTagIDs sets the recipe tags ids.
func SetTagIDs(ids []uuid.UUID) Option {
	return func(r *Recipe) error {
		r.tagIDs = ids
		return nil
	}
}

// Recipe is a cooking recipe.
type Recipe struct {
	id          uuid.UUID
	createdAt   time.Time
	updatedAt   *time.Time
	ownerID     uuid.UUID
	name        string
	desc        string
	url         *string
	servings    int
	ingredients []*Ingredient
	steps       []string
	tagIDs      []uuid.UUID
}

func (r *Recipe) loadOptions(opts ...Option) error {
	issues := []string{}

	for _, opt := range opts {
		if err := opt(r); err != nil {
			issues = append(issues, err.Error())
		}
	}

	if len(issues) != 0 {
		return &domain.ValidationError{Issues: issues}
	}

	return nil
}

// ID returns the recipe id.
func (r *Recipe) ID() uuid.UUID {
	return r.id
}

// CreatedAt returns a timestamp when the recipe is created.
func (r *Recipe) CreatedAt() time.Time {
	return r.createdAt
}

// UpdatedAt returns a timestamp when the recipe was last updated.
func (r *Recipe) UpdatedAt() *time.Time {
	return r.updatedAt
}

// Name returns the recipe name.
func (r *Recipe) Name() string {
	return r.name
}

// Description returns the recipe description.
func (r *Recipe) Description() string {
	return r.name
}

// URL returns the recipe source url.
func (r *Recipe) URL() *string {
	return r.url
}

// Servings returns the recipe serving count.
func (r *Recipe) Servings() int {
	return r.servings
}

// OwnerID returns the creator of the recipe.
func (r *Recipe) OwnerID() uuid.UUID {
	return r.ownerID
}

// Steps returns the recipe steps.
func (r *Recipe) Steps() []string {
	return r.steps
}

// Ingredients returns the recipe ingredients.
func (r *Recipe) Ingredients() []*Ingredient {
	return r.ingredients
}

// TagIDs returns the recipe tag ids.
func (r *Recipe) TagIDs() []uuid.UUID {
	return r.tagIDs
}

// Update an existing recipe.
func (r *Recipe) Update(timestamp time.Time, opts ...Option) error {
	now := timestamp.UTC()
	r.updatedAt = &now

	if err := r.loadOptions(opts...); err != nil {
		return err
	}

	return nil
}

// NewRecipe creates a new recipe.
func NewRecipe(id uuid.UUID, name string, ownerID uuid.UUID, timestamp time.Time, opts ...Option) (*Recipe, error) {
	recipe := &Recipe{
		id:          id,
		createdAt:   timestamp.UTC(),
		ownerID:     ownerID,
		name:        "",
		servings:    1,
		steps:       make([]string, 0),
		ingredients: make([]*Ingredient, 0),
		tagIDs:      make([]uuid.UUID, 0),
	}

	opts = append(opts, SetName(name))

	if err := recipe.loadOptions(opts...); err != nil {
		return nil, err
	}

	return recipe, nil
}
