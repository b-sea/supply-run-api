// Package unit defines everything to manage the units of measurement domain.
package unit

import "github.com/google/uuid"

// Repository defines all functions required to interact with units of measurement.
type Repository interface {
	Find(filter *Filter) ([]*Unit, error)
	Create(unit *Unit) error
	Update(unit *Unit) error
	Delete(id uuid.UUID) error
}

// Filter is a search filter for units.
type Filter struct {
	Owners []uuid.UUID
	System *System
	Type   *Type
}
