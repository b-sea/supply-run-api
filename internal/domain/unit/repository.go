package unit

import "github.com/google/uuid"

// Repository defines all functions required to interact with units of measurement.
type Repository interface {
	GetByOwnerID(id uuid.UUID) ([]*Unit, error)
	Create(unit *Unit) error
	Update(unit *Unit) error
	Delete(id uuid.UUID) error

	GetSystems() ([]*System, error)
	CreateSystem(system *System) error

	GetTypes() ([]*Type, error)
	CreateType(unitType *Type) error
}
