package unit

import (
	"errors"
	"fmt"
	"math"

	"github.com/b-sea/supply-run-api/internal/entity"
)

// Conversion defines a ratioed relationship between two units.
type Conversion struct {
	from  *Unit
	to    *Unit
	ratio float64
}

// NewConversion creates a new unit Conversion.
func NewConversion(from *Unit, to *Unit, ratio float64) (*Conversion, error) {
	validation := &entity.ValidationError{
		InnerErrors: make([]error, 0),
	}

	if from.ID() == to.ID() {
		validation.InnerErrors = append(validation.InnerErrors, errors.New("cannot convert to the same unit"))
	}

	if from.BaseType() != to.BaseType() {
		validation.InnerErrors = append(
			validation.InnerErrors,
			fmt.Errorf("cannot convert %q to %q", from.BaseType(), to.BaseType()),
		)
	}

	if !validation.IsEmpty() {
		return nil, validation
	}

	return &Conversion{
		from:  from,
		to:    to,
		ratio: ratio,
	}, nil
}

// From is the Unit to convert from.
func (c *Conversion) From() *Unit {
	return c.from
}

// To is the Unit to convert to.
func (c *Conversion) To() *Unit {
	return c.to
}

// Ratio is the amount of "to" units in a single "from" unit.
func (c *Conversion) Ratio() float64 {
	return c.ratio
}

// Kilo creates a new Unit 1000x larger than the given Unit and returns a Conversion between the two.
// The created unit will have a "kilo" prefix.
func Kilo(unit *Unit) *Conversion {
	return magnitude("kilo", "k", unit, 3) //nolint: mnd
}

// Centi creates a new Unit 100x larger than the given Unit and returns a Conversion between the two.
// The created unit will have a "centi" prefix.
func Centi(unit *Unit) *Conversion {
	return magnitude("centi", "c", unit, -2)
}

// Milli creates a new Unit 1000x smaller than the given Unit and returns a Conversion between the two.
// The created unit will have a "milli" prefix.
func Milli(unit *Unit) *Conversion {
	return magnitude("milli", "m", unit, -3)
}

func magnitude(prefix string, symbol string, unit *Unit, power float64, options ...Option) *Conversion {
	options = append(options, SetBaseType(unit.base), SetSystem(unit.system))

	return &Conversion{
		from:  unit,
		to:    New(prefix+unit.name, symbol+unit.symbol, options...),
		ratio: math.Pow(10, power), //nolint: mnd
	}
}
