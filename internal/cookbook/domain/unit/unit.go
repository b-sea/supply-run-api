// Package unit defines everything to manage the units of measurement domain.
package unit

import (
	"errors"
	"time"

	"github.com/b-sea/supply-run-api/internal/cookbook/domain"
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
	ownerID   uuid.UUID
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

// UpdatedAt returns a timestamp when the unit was last updated.
func (u *Unit) UpdatedAt() *time.Time {
	return u.updatedAt
}

// OwnerID returns the creator of the unit.
func (u *Unit) OwnerID() uuid.UUID {
	return u.ownerID
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
func (u *Unit) Update(timestamp time.Time, opts ...Option) error {
	now := timestamp.UTC()
	u.updatedAt = &now

	if err := u.loadOptions(opts...); err != nil {
		return err
	}

	return nil
}

// NewUnit creates a new unit.
func NewUnit(id uuid.UUID, name string, ownerID uuid.UUID, timestamp time.Time, opts ...Option) (*Unit, error) {
	unit := &Unit{
		id:        id,
		createdAt: timestamp.UTC(),
		ownerID:   ownerID,
		name:      "",
		siType:    NoType,
		system:    NoSystem,
	}

	opts = append(opts, SetName(name))

	if err := unit.loadOptions(opts...); err != nil {
		return nil, err
	}

	return unit, nil
}
