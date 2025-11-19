package unit_test

import (
	"testing"

	"github.com/b-sea/supply-run-api/internal/unit"
	"github.com/stretchr/testify/assert"
)

func TestNewConversion(t *testing.T) {
	t.Parallel()

	// Create a valid conversion
	from := unit.New("inch", "in")
	to := unit.New("foot", "ft")
	ratio := 12
	test, err := unit.NewConversion(from, to, float64(ratio))

	assert.NoError(t, err)
	assert.Equal(t, from, test.From())
	assert.Equal(t, to, test.To())
	assert.Equal(t, float64(ratio), test.Ratio())

	// Create a unit conversion to itself
	_, err = unit.NewConversion(from, from, float64(ratio))
	assert.Error(t, err)

	// Create a unit conversion to a unit of a different base type
	_, err = unit.NewConversion(from, unit.New("liter", "l", unit.Volume), float64(ratio))
	assert.Error(t, err)

}

func TestKilo(t *testing.T) {
	t.Parallel()

	gram := unit.New("gram", "g", unit.Metric, unit.Mass)
	test := unit.Kilo(gram)

	assert.Equal(t, gram, test.From())
	assert.Equal(t, "kilogram", test.To().Name())
	assert.Equal(t, "kg", test.To().Symbol())
	assert.Equal(t, "metric", test.To().System())
	assert.Equal(t, "mass", test.To().BaseType())
	assert.Equal(t, float64(1000), test.Ratio())
}

func TestCenti(t *testing.T) {
	t.Parallel()

	gram := unit.New("gram", "g", unit.Metric, unit.Mass)
	test := unit.Centi(gram)

	assert.Equal(t, gram, test.From())
	assert.Equal(t, "centigram", test.To().Name())
	assert.Equal(t, "cg", test.To().Symbol())
	assert.Equal(t, "metric", test.To().System())
	assert.Equal(t, "mass", test.To().BaseType())
	assert.Equal(t, float64(.01), test.Ratio())
}

func TestMilli(t *testing.T) {
	t.Parallel()

	gram := unit.New("gram", "g", unit.Metric, unit.Mass)
	test := unit.Milli(gram)

	assert.Equal(t, gram, test.From())
	assert.Equal(t, "milligram", test.To().Name())
	assert.Equal(t, "mg", test.To().Symbol())
	assert.Equal(t, "metric", test.To().System())
	assert.Equal(t, "mass", test.To().BaseType())
	assert.Equal(t, float64(.001), test.Ratio())
}
