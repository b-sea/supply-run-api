package unit_test

import (
	"testing"

	"github.com/b-sea/supply-run-api/internal/entity"
	"github.com/b-sea/supply-run-api/internal/unit"
	"github.com/stretchr/testify/assert"
)

func TestNewUnit(t *testing.T) {
	t.Parallel()

	// Create a valid unit
	name := "goober"
	symbol := "gr"
	test := unit.New(name, symbol)

	assert.Equal(t, entity.NewID(name), test.ID())
	assert.Equal(t, name, test.Name())
	assert.Equal(t, name+"s", test.Plural())
	assert.Equal(t, symbol, test.Symbol())
}
