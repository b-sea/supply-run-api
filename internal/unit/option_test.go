package unit_test

import (
	"testing"

	"github.com/b-sea/supply-run-api/internal/unit"
	"github.com/stretchr/testify/assert"
)

func TestSetSystem(t *testing.T) {
	t.Parallel()

	test := unit.New("goober", "gr")

	// Set the system
	unit.SetSystem("special")(test)
	assert.Equal(t, "special", test.System())
}

func TestSetBaseType(t *testing.T) {
	t.Parallel()

	test := unit.New("goober", "gr")

	// Set the base type
	unit.SetBaseType("diagonal")(test)
	assert.Equal(t, "diagonal", test.BaseType())
}

func TestWithCustomPlural(t *testing.T) {
	t.Parallel()

	test := unit.New("foot", "ft")

	// Set a custom plural
	unit.WithCustomPlural("feet")(test)
	assert.Equal(t, "feet", test.Plural())
}

func TestWithNoPlural(t *testing.T) {
	t.Parallel()

	test := unit.New("goober", "gr")

	// Remove plural
	unit.WithNoPlural()(test)
	assert.Equal(t, "goober", test.Plural())
}
