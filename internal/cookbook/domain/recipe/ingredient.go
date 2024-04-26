// Package recipe defines everything to manage the recipes domain.
package recipe

import (
	"errors"

	"github.com/b-sea/supply-run-api/internal/cookbook/domain"
	"github.com/google/uuid"
)

// IngredientOption is an ingredient creation option.
type IngredientOption func(i *Ingredient) error

// SetIngredientUnit sets the measurment unit of the ingredient.
func SetIngredientUnit(unitID uuid.UUID) IngredientOption {
	return func(i *Ingredient) error {
		i.unitID = unitID
		return nil
	}
}

// SetIngredientQuantity sets the quantity of the ingredient.
func SetIngredientQuantity(quantity float64) IngredientOption {
	return func(i *Ingredient) error {
		if quantity <= 0 {
			return errors.New("ingredient quantity must be greater than 0") //nolint: goerr113
		}

		i.quantity = quantity

		return nil
	}
}

// Ingredient is a recipe ingredient.
type Ingredient struct {
	id       uuid.UUID
	itemID   uuid.UUID
	unitID   uuid.UUID
	quantity float64
}

func (i *Ingredient) loadOptions(opts ...IngredientOption) error {
	issues := []string{}

	for _, opt := range opts {
		if err := opt(i); err != nil {
			issues = append(issues, err.Error())
		}
	}

	if len(issues) != 0 {
		return &domain.ValidationError{Issues: issues}
	}

	return nil
}

// ID returns the ingredient id.
func (i *Ingredient) ID() uuid.UUID {
	return i.id
}

// ItemID returns the item id used in the ingredient.
func (i *Ingredient) ItemID() uuid.UUID {
	return i.itemID
}

// UnitID returns the unit id used in the ingredient.
func (i *Ingredient) UnitID() uuid.UUID {
	return i.itemID
}

// Quantity returns the amount of ingredient used in the recipe.
func (i *Ingredient) Quantity() float64 {
	return i.quantity
}

// Update an ingredient.
func (i *Ingredient) Update(opts ...IngredientOption) error {
	if err := i.loadOptions(opts...); err != nil {
		return err
	}

	return nil
}

// NewIngredient creates a new recipe ingredient.
func NewIngredient(id uuid.UUID, itemID uuid.UUID, unitID uuid.UUID, quantity float64) (*Ingredient, error) {
	ingredient := &Ingredient{
		id:     id,
		itemID: itemID,
	}

	opts := []IngredientOption{SetIngredientUnit(unitID), SetIngredientQuantity(quantity)}
	if err := ingredient.loadOptions(opts...); err != nil {
		return nil, err
	}

	return ingredient, nil
}
