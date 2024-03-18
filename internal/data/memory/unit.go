// Package memory implements domains in memory storage.
package memory

import (
	"github.com/b-sea/supply-run-api/internal/domain/unit"
	"github.com/google/uuid"
)

// UnitRepository implements the unit domain repository.
type UnitRepository struct {
	units   map[uuid.UUID]*unit.Unit
	systems map[uuid.UUID]*unit.System
	types   map[uuid.UUID]*unit.Type
}

// NewUnitRepository creates a new UnitRepository.
func NewUnitRepository() *UnitRepository {
	return &UnitRepository{
		units:   make(map[uuid.UUID]*unit.Unit),
		systems: make(map[uuid.UUID]*unit.System),
		types:   make(map[uuid.UUID]*unit.Type),
	}
}

// GetByOwnerID finds all units based on the given owner.
func (r *UnitRepository) GetByOwnerID(id uuid.UUID) ([]*unit.Unit, error) {
	results := []*unit.Unit{}

	for k, v := range r.units {
		if k != id {
			continue
		}

		results = append(results, v)
	}

	return results, nil
}

// Create a new unit of measurement.
func (r *UnitRepository) Create(unit *unit.Unit) error {
	r.units[unit.ID()] = unit
	return nil
}

// Update an existing unit of measurement.
func (r *UnitRepository) Update(unit *unit.Unit) error {
	r.units[unit.ID()] = unit
	return nil
}

// Delete an existing unit of measurement.
func (r *UnitRepository) Delete(id uuid.UUID) error {
	delete(r.units, id)
	return nil
}

// GetSystems finds all measurement systems.
func (r *UnitRepository) GetSystems() ([]*unit.System, error) {
	results := []*unit.System{}

	for _, v := range r.systems {
		results = append(results, v)
	}

	return results, nil
}

// CreateSystem creates a new measurement system.
func (r *UnitRepository) CreateSystem(system *unit.System) error {
	r.systems[system.ID()] = system
	return nil
}

// GetTypes finds all SI unit types.
func (r *UnitRepository) GetTypes() ([]*unit.Type, error) {
	results := []*unit.Type{}

	for _, v := range r.types {
		results = append(results, v)
	}

	return results, nil
}

// CreateType creates a new SI unit type.
func (r *UnitRepository) CreateType(unitType *unit.Type) error {
	r.types[unitType.ID()] = unitType
	return nil
}
