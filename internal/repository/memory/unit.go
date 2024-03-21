// Package memory implements domains in memory storage.
package memory

// import (
// 	"slices"

// 	"github.com/b-sea/supply-run-api/internal/domain/unit"
// 	"github.com/google/uuid"
// )

// // UnitRepository implements the unit domain.
// type UnitRepository struct {
// 	units map[uuid.UUID]*unit.Unit
// }

// // NewUnitRepository creates a new UnitRepository.
// func NewUnitRepository(units []*unit.Unit) *UnitRepository {
// 	result := UnitRepository{
// 		units: make(map[uuid.UUID]*unit.Unit),
// 	}

// 	for _, u := range units {
// 		result.units[u.ID()] = u
// 	}

// 	return &result
// }

// func (r *UnitRepository) isValid(entity *unit.Unit, filter *unit.Filter) bool {
// 	if filter == nil {
// 		return true
// 	}

// 	if filter.Owners != nil {
// 		if !slices.Contains(filter.Owners, entity.Owner()) {
// 			return false
// 		}
// 	}

// 	if filter.System != nil {
// 		return filter.System == &entity.System
// 	}

// 	if filter.Type != nil {
// 		return filter.Type == &entity.Type
// 	}

// 	return true
// }

// // Find searches for units.
// func (r *UnitRepository) Find(filter *unit.Filter) ([]*unit.Unit, error) {
// 	results := []*unit.Unit{}

// 	for _, entity := range r.units {
// 		if !r.isValid(entity, filter) {
// 			continue
// 		}

// 		results = append(results, entity)
// 	}

// 	return results, nil
// }

// // Create a new unit of measurement.
// func (r *UnitRepository) Create(unit *unit.Unit) error {
// 	r.units[unit.ID()] = unit
// 	return nil
// }

// // Update an existing unit of measurement.
// func (r *UnitRepository) Update(unit *unit.Unit) error {
// 	r.units[unit.ID()] = unit
// 	return nil
// }

// // Delete an existing unit of measurement.
// func (r *UnitRepository) Delete(id uuid.UUID) error {
// 	delete(r.units, id)
// 	return nil
// }
