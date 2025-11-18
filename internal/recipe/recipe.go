// Package recipe defines cookbook recipes.
package recipe

import (
	"time"

	"github.com/b-sea/supply-run-api/internal/entity"
)

// Recipe is a cookbook recipe.
type Recipe struct {
	id          entity.ID
	name        string
	url         string
	numServings int
	steps       []string
	ingredients []Ingredient
	tags        []string

	createdAt time.Time
	createdBy entity.ID
	updatedAt time.Time
	updatedBy entity.ID
}

// New creates a new Recipe.
func New(id entity.ID, name string, timestamp time.Time, userID entity.ID, options ...Option) (*Recipe, error) {
	recipe := &Recipe{
		id:          id,
		name:        "",
		numServings: 0,
		steps:       make([]string, 0),
		ingredients: make([]Ingredient, 0),
		tags:        make([]string, 0),
		createdAt:   timestamp.UTC(),
		createdBy:   userID,
		updatedAt:   timestamp.UTC(),
		updatedBy:   userID,
	}

	options = append([]Option{SetName(name), SetNumServings(recipe.numServings)}, options...)

	validation := &entity.ValidationError{
		InnerErrors: make([]error, 0),
	}

	for _, option := range options {
		if _, err := option(recipe); err != nil {
			validation.InnerErrors = append(validation.InnerErrors, err)

			continue
		}
	}

	if !validation.IsEmpty() {
		return nil, validation
	}

	return recipe, nil
}

// Update updates an existing Recipe.
// Updates are only applied if any data actually changes.
func (r *Recipe) Update(timestamp time.Time, userID entity.ID, options ...Option) error {
	changed := false
	validation := &entity.ValidationError{
		InnerErrors: make([]error, 0),
	}

	for _, option := range options {
		result, err := option(r)
		if err != nil {
			validation.InnerErrors = append(validation.InnerErrors, err)

			continue
		}

		if !result {
			continue
		}

		changed = true
	}

	if !validation.IsEmpty() {
		return validation
	}

	if !changed {
		return nil
	}

	r.updatedAt = timestamp.UTC()
	r.updatedBy = userID

	return nil
}

// ID returns the Recipe id.
func (r *Recipe) ID() entity.ID {
	return r.id
}

// Name returns the Recipe name.
func (r *Recipe) Name() string {
	return r.name
}

// URL returns the Recipe source url.
func (r *Recipe) URL() string {
	return r.url
}

// NumServings returns the number of servings in the Recipe.
func (r *Recipe) NumServings() int {
	return r.numServings
}

// Steps returns the Recipe steps.
func (r *Recipe) Steps() []string {
	return r.steps
}

// Ingredients returns the Recipe ingredients.
func (r *Recipe) Ingredients() []Ingredient {
	return r.ingredients
}

// Tags returns the Recipe tags.
func (r *Recipe) Tags() []string {
	return r.tags
}

// CreatedAt returns when the Recipe was created.
func (r *Recipe) CreatedAt() time.Time {
	return r.createdAt
}

// CreatedBy returns the id of the user who created the Recipe.
func (r *Recipe) CreatedBy() entity.ID {
	return r.createdBy
}

// UpdatedAt returns when the Recipe was last updated.
func (r *Recipe) UpdatedAt() time.Time {
	return r.updatedAt
}

// UpdatedBy returns the id of the user who last updated the Recipe.
func (r *Recipe) UpdatedBy() entity.ID {
	return r.updatedBy
}
