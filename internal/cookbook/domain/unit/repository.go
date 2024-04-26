// Package unit defines everything to manage the units of measurement domain.
package unit

import "github.com/google/uuid"

type Repository interface {
	Create(unit *Unit) error
	Delete(id uuid.UUID) error
}
