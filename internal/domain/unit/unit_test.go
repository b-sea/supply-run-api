package unit_test

import (
	"testing"
	"time"

	"github.com/b-sea/supply-run-api/internal/domain/unit"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUnitNewUnit(t *testing.T) {
	t.Parallel()

	id := uuid.New()
	owner := uuid.New()
	now := time.Now()

	testUnit := unit.NewUnit(
		"custom unit",
		owner,
		unit.WithID(id),
		unit.WithTimestamp(now),
		unit.WithSymbol("cu"),
		unit.WithSystem(unit.MetricSystem),
		unit.WithType(unit.VolumeType),
	)

	assert.Equal(t, id, testUnit.ID())
	assert.Equal(t, now, testUnit.CreatedAt())
	assert.Equal(t, owner, testUnit.Owner())
	assert.Equal(t, "custom unit", testUnit.Name)
	assert.Equal(t, "cu", testUnit.Symbol)
	assert.Equal(t, unit.VolumeType, testUnit.Type)
	assert.Equal(t, unit.MetricSystem, testUnit.System)
}

func TestUnitValidate(t *testing.T) {
	t.Parallel()

	testUnit := unit.NewUnit("", uuid.New())
	err := testUnit.Validate()

	assert.EqualError(t, err, "validation errors: name cannot be empty")

	testUnit = unit.NewUnit("custom unit", uuid.New())
	err = testUnit.Validate()

	assert.NoError(t, err)
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
