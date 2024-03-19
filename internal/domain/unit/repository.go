// Package unit defines everything to manage the units of measurement domain.
package unit

import "github.com/google/uuid"

// Repository defines all functions required to interact with units of measurement.
type Repository interface {
	GetByOwnerIDs(ids []uuid.UUID) ([]*Unit, error)
	Create(unit *Unit) error
	Update(unit *Unit) error
	Delete(id uuid.UUID) error
}
