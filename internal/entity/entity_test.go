package entity_test

import (
	"errors"
	"testing"

	"github.com/b-sea/supply-run-api/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewSeededID(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "hanuPt54naGV7gidHVi2NY", entity.NewSeededID([]byte("some-id-seed")).String())
	assert.Equal(t, "hanuPt54naGV7gidHVi2NY", entity.NewSeededID([]byte("some-id-seed")).String())

	assert.Equal(t, "CqdWFwQjLtbU2qM4r7umBc", entity.NewSeededID([]byte("a different seed")).String())
	assert.Equal(t, "CqdWFwQjLtbU2qM4r7umBc", entity.NewSeededID([]byte("a different seed")).String())
}

func TestValidationError(t *testing.T) {
	t.Parallel()

	err := entity.ValidationError{
		InnerErrors: []error{
			errors.New("oh no"),
			errors.New("something else"),
		},
	}
	assert.Equal(t, "validation errors: oh no, something else", err.Error())
	assert.False(t, err.IsEmpty())

	err = entity.ValidationError{
		InnerErrors: []error{},
	}
	assert.True(t, err.IsEmpty())
}
