// Package unit defines everything to manage the units of measurement domain.
package unit

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

type (
	// UUIDFunc is a function to generate an uuid.
	UUIDFunc func() uuid.UUID

	// TimeFunc is a function to generate a timestamp.
	TimeFunc func() time.Time
)

// ValidationError is raised when validation fails.
type ValidationError struct {
	Issues []string
}

func (e *ValidationError) Error() string {
	return "validation errors: " + strings.Join(e.Issues, ", ")
}

// Conversion is a ratioed relationship between units.
type Conversion struct {
	unit  Unit
	ratio float32
}

// ConversionInput is used to build a ratioed relationship between units.
type ConversionInput struct {
	Unit  Unit
	Ratio float32
}

// Unit is a unit of measurement.
type Unit struct {
	id          uuid.UUID
	createdAt   time.Time
	updatedAt   *time.Time
	owner       uuid.UUID
	name        string
	symbol      string
	siType      Type
	system      System
	convertFrom *Conversion
}

// ID returns the id of a unit.
func (u *Unit) ID() uuid.UUID {
	return u.id
}

// CreatedAt returns a timestamp when the unit is created.
func (u *Unit) CreatedAt() time.Time {
	return u.createdAt
}

// UpdatedAt returns a timestamp when the unit was updated.
func (u *Unit) UpdatedAt() *time.Time {
	return u.updatedAt
}

// Owner returns the creator of the unit.
func (u *Unit) Owner() uuid.UUID {
	return u.owner
}

// Name returns the unit name.
func (u *Unit) Name() string {
	return u.name
}

// Symbol returns the unit symbol/shortname.
func (u *Unit) Symbol() string {
	return u.symbol
}

// Type returns the unit SI type.
func (u *Unit) Type() Type {
	return u.siType
}

// System returns the unit measurement system.
func (u *Unit) System() System {
	return u.system
}

func (u *Unit) validate() error {
	issues := []string{}

	if u.name == "" {
		issues = append(issues, "name cannot be empty")
	}

	if u.convertFrom != nil {
		if u.id == u.convertFrom.unit.id {
			issues = append(issues, "cannot convert from self")
		}
	}

	if len(issues) > 0 {
		return &ValidationError{
			Issues: issues,
		}
	}

	return nil
}

// UpdateUnitInput is used to update an existing unit.
type UpdateUnitInput struct {
	Now          TimeFunc
	Name         *string
	Symbol       *string
	Type         *Type
	System       *System
	ConvertFrom  *Conversion
	ClearConvert bool
}

// Update and validate an existing unit.
func (u *Unit) Update(input UpdateUnitInput) error {
	if input.Now == nil {
		input.Now = time.Now
	}

	now := input.Now().UTC()
	u.updatedAt = &now

	if input.Name != nil {
		u.name = *input.Name
	}

	if input.Symbol != nil {
		u.symbol = *input.Symbol
	}

	if input.System != nil {
		u.symbol = *input.Symbol
	}

	if input.Type != nil {
		u.siType = *input.Type
	}

	if input.ConvertFrom != nil {
		u.convertFrom = input.ConvertFrom
	}

	if input.ClearConvert {
		u.convertFrom = nil
	}

	if err := u.validate(); err != nil {
		return err
	}

	return nil
}

// NewUnitInput is used to create a new unit.
type NewUnitInput struct {
	ID          UUIDFunc
	Now         TimeFunc
	Owner       uuid.UUID
	Name        string
	Symbol      string
	Type        Type
	System      System
	ConvertFrom *ConversionInput
}

// NewUnit creates and validates a new unit.
func NewUnit(input NewUnitInput) (*Unit, error) {
	if input.ID == nil {
		input.ID = uuid.New
	}

	if input.Now == nil {
		input.Now = time.Now
	}

	result := Unit{
		id:        input.ID(),
		createdAt: input.Now().UTC(),
		owner:     input.Owner,
		name:      input.Name,
		siType:    input.Type,
		symbol:    input.Symbol,
		system:    input.System,
	}

	if input.ConvertFrom != nil {
		result.convertFrom = &Conversion{
			unit:  input.ConvertFrom.Unit,
			ratio: input.ConvertFrom.Ratio,
		}
	}

	if err := result.validate(); err != nil {
		return nil, err
	}

	return &result, nil
}
