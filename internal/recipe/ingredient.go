package recipe

import units "github.com/bcicen/go-units"

// Ingredient is a Recipe ingredient.
type Ingredient struct {
	name  string
	value units.Value
}

// Name returns the Ingredient name.
func (i *Ingredient) Name() string {
	return i.name
}

// Quantity returns the amount of Ingredient.
func (i *Ingredient) Quantity() float64 {
	return i.value.Float()
}

// Unit returns the Ingredient default unit.
func (i *Ingredient) Unit() units.Unit {
	return i.value.Unit()
}
