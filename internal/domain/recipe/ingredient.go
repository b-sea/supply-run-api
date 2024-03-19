// Package recipe defines everything to manage the recipes domain.
package recipe

import (
	"github.com/b-sea/supply-run-api/internal/entity"
	"github.com/google/uuid"
)

// IngredientOption is an ingredient creation option.
type IngredientOption func(*Ingredient)

// WithIngredientID sets the ingredient id.
func WithIngredientID(id uuid.UUID) IngredientOption {
	return func(i *Ingredient) {
		i.id = id
	}
}

// Ingredient is a recipe ingredient.
type Ingredient struct {
	id       uuid.UUID
	item     uuid.UUID
	Unit     uuid.UUID
	Quantity float64
}

// ID returns the ingredient id.
func (i *Ingredient) ID() uuid.UUID {
	return i.id
}

// Item returns the item used in the ingredient.
func (i *Ingredient) Item() uuid.UUID {
	return i.item
}

// Validate the ingredient.
func (i *Ingredient) Validate() error {
	issues := []string{}

	if i.Quantity <= 0 {
		issues = append(issues, "quantity cannot be 0 or less than 0")
	}

	if len(issues) == 0 {
		return nil
	}

	return &entity.ValidationError{
		Issues: issues,
	}
}

// NewIngredient creates a new recipe ingredient.
func NewIngredient(item uuid.UUID, unit uuid.UUID, quantity float64, opts ...IngredientOption) *Ingredient {
	ingredient := &Ingredient{
		id:       uuid.New(),
		item:     item,
		Unit:     unit,
		Quantity: quantity,
	}

	for _, opt := range opts {
		opt(ingredient)
	}

	return ingredient
}
