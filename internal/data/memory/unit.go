package memory

import (
	"github.com/b-sea/supply-run-api/internal/domain/unit"
	"github.com/google/uuid"
)

type UnitRepository struct {
	units   map[uuid.UUID]*unit.Unit
	systems map[uuid.UUID]*unit.System
	types   map[uuid.UUID]*unit.Type
}

func NewUnitRepository() *UnitRepository {
	return &UnitRepository{
		units:   make(map[uuid.UUID]*unit.Unit),
		systems: make(map[uuid.UUID]*unit.System),
		types:   make(map[uuid.UUID]*unit.Type),
	}
}

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

func (r *UnitRepository) Create(unit *unit.Unit) error {
	r.units[unit.ID()] = unit
	return nil
}

func (r *UnitRepository) Update(unit *unit.Unit) error {
	r.units[unit.ID()] = unit
	return nil
}

func (r *UnitRepository) Delete(id uuid.UUID) error {
	delete(r.units, id)
	return nil
}

func (r *UnitRepository) GetSystems() ([]*unit.System, error) {
	results := []*unit.System{}

	for _, v := range r.systems {
		results = append(results, v)
	}

	return results, nil
}

func (r *UnitRepository) CreateSystem(system *unit.System) error {
	r.systems[system.ID()] = system
	return nil
}

func (r *UnitRepository) GetTypes() ([]*unit.Type, error) {
	results := []*unit.Type{}

	for _, v := range r.types {
		results = append(results, v)
	}

	return results, nil
}

func (r *UnitRepository) CreateType(unitType *unit.Type) error {
	r.types[unitType.ID()] = unitType
	return nil
}
