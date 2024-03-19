// Package memory implements domains in memory storage.
package memory

import (
	"slices"

	"github.com/b-sea/supply-run-api/internal/domain/unit"
	"github.com/google/uuid"
)

// UnitRepository implements the unit domain.
type UnitRepository struct {
	units map[uuid.UUID]*unit.Unit
}

// NewUnitRepository creates a new UnitRepository.
func NewUnitRepository(units []*unit.Unit) *UnitRepository {
	result := UnitRepository{
		units: make(map[uuid.UUID]*unit.Unit),
	}

	for _, u := range units {
		result.units[u.ID()] = u
	}

	return &result
}

// GetByOwnerIDs finds all units based on the given owner.
func (r *UnitRepository) GetByOwnerIDs(ids []uuid.UUID) ([]*unit.Unit, error) {
	results := []*unit.Unit{}

	for _, v := range r.units {
		if !slices.Contains(ids, v.Owner()) {
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
