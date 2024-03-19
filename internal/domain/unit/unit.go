// Package unit defines everything to manage the units of measurement domain.
package unit

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

// Type is a base or derived SI measurment types.
type Type int

// NoType et all are the different SI types.
const (
	NoType Type = iota
	MassType
	VolumeType
)

func (t Type) String() string {
	switch t {
	case MassType:
		return "MASS"
	case VolumeType:
		return "VOLUME"
	case NoType:
		fallthrough
	default:
		return ""
	}
}

// TypeFromString converts a string to a SI type.
func TypeFromString(s string) Type {
	switch s {
	case "MASS":
		return MassType
	case "VOLUME":
		return VolumeType
	default:
		return NoType
	}
}

// System is a measurement system.
type System int

// NoSystem et all are the different measurement systems.
const (
	NoSystem System = iota
	ImperialSystem
	MetricSystem
)

func (s System) String() string {
	switch s {
	case ImperialSystem:
		return "IMPERIAL"
	case MetricSystem:
		return "METRIC"
	case NoSystem:
		fallthrough
	default:
		return ""
	}
}

// SystemFromString converts a string to a measurement system.
func SystemFromString(s string) System {
	switch s {
	case "IMPERIAL":
		return ImperialSystem
	case "METRIC":
		return MetricSystem
	default:
		return NoSystem
	}
}

// ValidationError is raised when validation fails.
type ValidationError struct {
	Issues []string
}

func (e *ValidationError) Error() string {
	return "validation errors: " + strings.Join(e.Issues, ", ")
}

// Unit is a unit of measurement.
type Unit struct {
	id        uuid.UUID
	createdAt time.Time
	owner     uuid.UUID
	Name      string
	Symbol    string
	Type      Type
	System    System
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

// Validate the unit.
func (u *Unit) Validate() error {
	issues := []string{}

	if u.Name == "" {
		issues = append(issues, "name cannot be empty")
	}

	if len(issues) == 0 {
		return nil
	}

	return &ValidationError{
		Issues: issues,
	}
}

// Option is a unit creation option.
type Option func(*Unit)

// WithID sets the unit id.
func WithID(id uuid.UUID) Option {
	return func(u *Unit) {
		u.id = id
	}
}

// WithTimestamp sets the unit creation time.
func WithTimestamp(now time.Time) Option {
	return func(u *Unit) {
		u.createdAt = now
	}
}

// WithSymbol sets the unit symbol.
func WithSymbol(symbol string) Option {
	return func(u *Unit) {
		u.Symbol = symbol
	}
}

// WithSystem sets the unit measurment system.
func WithSystem(system System) Option {
	return func(u *Unit) {
		u.System = system
	}
}

// WithType sets the unit SI type.
func WithType(siType Type) Option {
	return func(u *Unit) {
		u.Type = siType
	}
}

// NewUnit creates a new unit.
func NewUnit(name string, owner uuid.UUID, opts ...Option) *Unit {
	result := &Unit{
		id:        uuid.New(),
		createdAt: time.Now().UTC(),
		owner:     owner,
		Name:      name,
		Symbol:    "",
		Type:      NoType,
		System:    NoSystem,
	}

	for _, opt := range opts {
		opt(result)
	}

	return result
}
