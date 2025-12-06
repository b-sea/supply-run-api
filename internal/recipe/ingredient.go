package recipe

import (
	"github.com/b-sea/supply-run-api/internal/entity"
)

// Ingredient is a Recipe ingredient.
type Ingredient struct {
	name     string
	quantity float64
	unitID   entity.ID
}

// Name returns the Ingredient name.
func (i *Ingredient) Name() string {
	return i.name
}

// Quantity returns the amount of Ingredient.
func (i *Ingredient) Quantity() float64 {
	return i.quantity
}

// UnitID returns the Ingredient unit id.
func (i *Ingredient) UnitID() entity.ID {
	return i.unitID
}
