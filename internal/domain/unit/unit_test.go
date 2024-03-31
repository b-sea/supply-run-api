package unit_test

import (
	"testing"
	"time"

	"github.com/b-sea/supply-run-api/internal/domain"
	"github.com/b-sea/supply-run-api/internal/domain/unit"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUnitNewUnit(t *testing.T) {
	t.Parallel()

	type test struct {
		id        uuid.UUID
		timestamp time.Time
		ownerID   uuid.UUID
		name      string
		err       error
	}

	testCases := map[string]test{
		"success": {
			name: "special unit",
		},
		"invalid": {
			name: "",
			err:  &domain.ValidationError{Issues: []string{"unit name cannot be empty"}},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(s *testing.T) {
			s.Parallel()

			newUnit, err := unit.NewUnit(testCase.id, testCase.timestamp, testCase.name, testCase.ownerID)
			if err == nil {
				assert.Equal(t, testCase.id, newUnit.ID(), "different ID")
				assert.Equal(t, testCase.timestamp, newUnit.CreatedAt(), "different CreatedAt")
				assert.Equal(t, testCase.ownerID, newUnit.OwnerID(), "different OwnerID")
				assert.Equal(t, testCase.name, newUnit.Name(), "different Name")
			}

			if testCase.err == nil {
				assert.NoError(t, err, "no error expected")
			} else {
				assert.EqualError(t, err, testCase.err.Error(), "different errors")
			}
		})
	}
}

func TestUnitUpdate(t *testing.T) {
	t.Parallel()

	newUnit, err := unit.NewUnit(uuid.New(), time.Now(), "custom unit", uuid.New())
	if err != nil {
		panic("creating test unit failed")
	}

	type test struct {
		timestamp time.Time
		err       error
	}

	testCases := map[string]test{
		"success": {
			timestamp: time.Now(),
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(s *testing.T) {
			s.Parallel()

			err := newUnit.Update(testCase.timestamp)
			if err == nil {
				updatedAt := testCase.timestamp.UTC()
				assert.Equal(t, &updatedAt, newUnit.UpdatedAt(), "different results")
			}

			if testCase.err == nil {
				assert.NoError(t, err, "no error expected")
			} else {
				assert.EqualError(t, err, testCase.err.Error(), "different errors")
			}
		})
	}
}

func TestUnitSetName(t *testing.T) {
	t.Parallel()

	newUnit, err := unit.NewUnit(uuid.New(), time.Now(), "custom unit", uuid.New())
	if err != nil {
		panic("creating test unit failed")
	}

	type test struct {
		name string
		err  error
	}

	testCases := map[string]test{
		"success": {
			name: "special unit",
		},
		"invalid": {
			name: "",
			err:  &domain.ValidationError{Issues: []string{"unit name cannot be empty"}},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(s *testing.T) {
			s.Parallel()

			err := newUnit.Update(time.Now(), unit.SetName(testCase.name))
			if err == nil {
				assert.Equal(t, testCase.name, newUnit.Name(), "different results")
			}

			if testCase.err == nil {
				assert.NoError(t, err, "no error expected")
			} else {
				assert.EqualError(t, err, testCase.err.Error(), "different errors")
			}
		})
	}
}

func TestUnitSetSymbol(t *testing.T) {
	t.Parallel()

	newUnit, err := unit.NewUnit(uuid.New(), time.Now(), "custom unit", uuid.New())
	if err != nil {
		panic("creating test unit failed")
	}

	type test struct {
		symbol string
		err    error
	}

	testCases := map[string]test{
		"success": {
			symbol: "s",
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(s *testing.T) {
			s.Parallel()

			err := newUnit.Update(time.Now(), unit.SetSymbol(testCase.symbol))
			if err == nil {
				assert.Equal(t, testCase.symbol, newUnit.Symbol(), "different results")
			}

			if testCase.err == nil {
				assert.NoError(t, err, "no error expected")
			} else {
				assert.EqualError(t, err, testCase.err.Error(), "different errors")
			}
		})
	}
}

func TestUnitSetType(t *testing.T) {
	t.Parallel()

	newUnit, err := unit.NewUnit(uuid.New(), time.Now(), "custom unit", uuid.New())
	if err != nil {
		panic("creating test unit failed")
	}

	type test struct {
		siType unit.Type
		err    error
	}

	testCases := map[string]test{
		"success": {
			siType: unit.MassType,
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(s *testing.T) {
			s.Parallel()

			err := newUnit.Update(time.Now(), unit.SetType(testCase.siType))
			if err == nil {
				assert.Equal(t, testCase.siType, newUnit.Type(), "different results")
			}

			if testCase.err == nil {
				assert.NoError(t, err, "no error expected")
			} else {
				assert.EqualError(t, err, testCase.err.Error(), "different errors")
			}
		})
	}
}

func TestUnitSetSystem(t *testing.T) {
	t.Parallel()

	newUnit, err := unit.NewUnit(uuid.New(), time.Now(), "custom unit", uuid.New())
	if err != nil {
		panic("creating test unit failed")
	}

	type test struct {
		system unit.System
		err    error
	}

	testCases := map[string]test{
		"success": {
			system: unit.ImperialSystem,
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(s *testing.T) {
			s.Parallel()

			err := newUnit.Update(time.Now(), unit.SetSystem(testCase.system))
			if err == nil {
				assert.Equal(t, testCase.system, newUnit.System(), "different results")
			}

			if testCase.err == nil {
				assert.NoError(t, err, "no error expected")
			} else {
				assert.EqualError(t, err, testCase.err.Error(), "different errors")
			}
		})
	}
}

func TestUnitEnum(t *testing.T) {
	t.Parallel()

	assert.Equal(t, unit.ImperialSystem, unit.SystemFromString("IMPERIAL"))
	assert.Equal(t, "IMPERIAL", unit.ImperialSystem.String())
	assert.Equal(t, unit.MetricSystem, unit.SystemFromString("METRIC"))
	assert.Equal(t, "METRIC", unit.MetricSystem.String())
	assert.Equal(t, unit.NoSystem, unit.SystemFromString("SOMETHING ELSE"))
	assert.Equal(t, "", unit.NoSystem.String())

	assert.Equal(t, unit.MassType, unit.TypeFromString("MASS"))
	assert.Equal(t, "MASS", unit.MassType.String())
	assert.Equal(t, unit.VolumeType, unit.TypeFromString("VOLUME"))
	assert.Equal(t, "VOLUME", unit.VolumeType.String())
	assert.Equal(t, unit.NoType, unit.TypeFromString("SOMETHING ELSE"))
	assert.Equal(t, "", unit.NoType.String())
}
