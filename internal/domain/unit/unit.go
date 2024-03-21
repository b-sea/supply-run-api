// Package unit defines everything to manage the units of measurement domain.
package unit

import (
	"errors"
	"time"

	"github.com/b-sea/supply-run-api/internal/domain"
	"github.com/google/uuid"
)

// Option is a unit creation option.
type Option func(*Unit) error

// SetName sets the unit name.
func SetName(name string) Option {
	return func(u *Unit) error {
		if name == "" {
			return errors.New("unit name cannot be empty") //nolint: goerr113
		}

		u.name = name

		return nil
	}
}

// SetSymbol sets the unit symbol.
func SetSymbol(symbol string) Option {
	return func(u *Unit) error {
		u.symbol = symbol
		return nil
	}
}

// SetSystem sets the unit measurment system.
func SetSystem(system System) Option {
	return func(u *Unit) error {
		u.system = system
		return nil
	}
}

// SetType sets the unit SI type.
func SetType(siType Type) Option {
	return func(u *Unit) error {
		u.siType = siType
		return nil
	}
}

// Unit is a unit of measurement.
type Unit struct {
	id        uuid.UUID
	createdAt time.Time
	updatedAt *time.Time
	owner     uuid.UUID
	name      string
	symbol    string
	siType    Type
	system    System
}

func (u *Unit) loadOptions(opts ...Option) error {
	issues := []string{}

	for _, opt := range opts {
		if err := opt(u); err != nil {
			issues = append(issues, err.Error())
		}
	}

	if len(issues) != 0 {
		return &domain.ValidationError{Issues: issues}
	}

	return nil
}

// ID returns the id of a unit.
func (u *Unit) ID() uuid.UUID {
	return u.id
}

// CreatedAt returns a timestamp when the unit is created.
func (u *Unit) CreatedAt() time.Time {
	return u.createdAt
}

// Owner returns the creator of the unit.
func (u *Unit) Owner() uuid.UUID {
	return u.owner
}

// Name returns the name of the unit.
func (u *Unit) Name() string {
	return u.name
}

// Symbol returns the symbol of the unit.
func (u *Unit) Symbol() string {
	return u.symbol
}

// Type returns the SI type of the unit.
func (u *Unit) Type() Type {
	return u.siType
}

// System returns the measurement system of the unit.
func (u *Unit) System() System {
	return u.system
}

// Update an existing unit.
func (u *Unit) Update(opts ...Option) error {
	now := time.Now().UTC()
	u.updatedAt = &now

	if err := u.loadOptions(opts...); err != nil {
		return err
	}

	return nil
}

// NewUnit creates a new unit.
func NewUnit(name string, owner uuid.UUID, opts ...Option) (*Unit, error) {
	unit := &Unit{
		id:        uuid.New(),
		createdAt: time.Now().UTC(),
		owner:     owner,
		name:      name,
		siType:    NoType,
		system:    NoSystem,
	}

	if err := unit.loadOptions(opts...); err != nil {
		return nil, err
	}

	return unit, nil
}

// Hydrate returns a unit in an existing state.
func Hydrate(id uuid.UUID, name string, createdAt time.Time, updatedAt *time.Time, ownerID uuid.UUID) (*Unit, error) {
	unit, err := NewUnit(name, ownerID)
	if err != nil {
		return nil, err
	}

	unit.id = id
	unit.createdAt = createdAt
	unit.updatedAt = updatedAt

	return unit, nil
}
