// Package unit defines units of measurement.
package unit

import (
	"github.com/b-sea/supply-run-api/internal/entity"
)

// Unit is a unit of measurement.
type Unit struct {
	id     entity.ID
	name   string
	plural string
	symbol string
	base   string
	system string
}

// New creates a new Unit.
func New(name string, symbol string, options ...Option) *Unit {
	unit := &Unit{
		name:   name,
		plural: name + "s",
		symbol: symbol,
	}

	for _, option := range options {
		option(unit)
	}

	unit.id = entity.NewID(unit.system + unit.base + name)

	return unit
}

// ID returns the Unit id.
// Unit IDs are deterministic, based off the unit's name, system, and base.
func (u *Unit) ID() entity.ID {
	return u.id
}

// Name returns the Unit name.
func (u *Unit) Name() string {
	return u.name
}

// Plural returns the Unit plural name.
func (u *Unit) Plural() string {
	return u.plural
}

// Symbol returns the Unit symbol.
func (u *Unit) Symbol() string {
	return u.symbol
}

// BaseType returns the base type of unit (mass, length, temperature, etc).
func (u *Unit) BaseType() string {
	return u.base
}

// System returns the Unit system (metric, imperial, us, etc).
func (u *Unit) System() string {
	return u.system
}
