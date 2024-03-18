package unit_test

import (
	"testing"
	"time"

	"github.com/b-sea/supply-run-api/internal/domain/unit"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// func TestUnitCreate(t *testing.T) {
// 	t.Parallel()

// 	admin := uuid.New()

// 	imperial, err := unit.NewSystem(
// 		unit.NewSystemInput{
// 			Owner: admin,
// 			Name:  "imperial",
// 		},
// 	)
// 	if err != nil {
// 		panic(err)
// 	}
// 	logrus.Infof("%+v", imperial)

// 	metric, err := unit.NewSystem(
// 		unit.NewSystemInput{
// 			Owner: admin,
// 			Name:  "metric",
// 		},
// 	)
// 	if err != nil {
// 		panic(err)
// 	}
// 	logrus.Infof("%+v", metric)

// 	volume, err := unit.NewType(
// 		unit.NewTypeInput{
// 			Owner: admin,
// 			Name:  "volume",
// 		},
// 	)
// 	if err != nil {
// 		panic(err)
// 	}
// 	logrus.Infof("%+v", volume)

// 	mass, err := unit.NewType(
// 		unit.NewTypeInput{
// 			Owner: admin,
// 			Name:  "mass",
// 		},
// 	)
// 	if err != nil {
// 		panic(err)
// 	}
// 	logrus.Infof("%+v", mass)

// 	ml, err := unit.NewUnit(
// 		unit.NewUnitInput{
// 			Owner:  admin,
// 			Name:   "millilitre",
// 			Symbol: "ml",
// 			System: *metric,
// 			Type:   *volume,
// 		},
// 	)
// 	if err != nil {
// 		panic(err)
// 	}
// 	logrus.Infof("%+v", ml)

// 	tsp, err := unit.NewUnit(
// 		unit.NewUnitInput{
// 			Owner:  admin,
// 			Name:   "teaspoon",
// 			Symbol: "tsp",
// 			System: *imperial,
// 			Type:   *volume,
// 			ConvertFrom: &unit.ConversionInput{
// 				Unit:  *ml,
// 				Ratio: 4.929,
// 			},
// 		},
// 	)
// 	if err != nil {
// 		panic(err)
// 	}
// 	logrus.Infof("%+v", tsp)

// 	tbsp, err := unit.NewUnit(
// 		unit.NewUnitInput{
// 			Owner:  admin,
// 			Name:   "tablespoon",
// 			Symbol: "tbsp",
// 			System: *imperial,
// 			Type:   *volume,
// 			ConvertFrom: &unit.ConversionInput{
// 				Unit:  *tsp,
// 				Ratio: 2,
// 			},
// 		},
// 	)
// 	if err != nil {
// 		panic(err)
// 	}
// 	logrus.Infof("%+v", tbsp)
// }

func TestUnitNewUnit(t *testing.T) {
	t.Parallel()

	generated_id := uuid.New()
	idFunc := func() uuid.UUID {
		return generated_id
	}
	timestamp := time.Now
	owner := uuid.New()

	newSystem, _ := unit.NewSystem(
		unit.NewSystemInput{
			Owner: owner,
			Name:  "imperial",
		},
	)

	newType, _ := unit.NewType(
		unit.NewTypeInput{
			Owner: owner,
			Name:  "mass",
		},
	)

	newUnit, err := unit.NewUnit(
		unit.NewUnitInput{
			ID:     idFunc,
			Now:    timestamp,
			Owner:  owner,
			Name:   "firkin",
			Symbol: "fk",
			Type:   *newType,
			System: *newSystem,
		},
	)

	assert.NoError(t, err)
	assert.Equal(t, idFunc(), newUnit.ID())
	assert.Equal(t, timestamp().UTC(), newUnit.CreatedAt())
	assert.Nil(t, newUnit.UpdatedAt())
	assert.Equal(t, owner, newUnit.Owner())
	assert.Equal(t, "firkin", newUnit.Name())
	assert.Equal(t, "fk", newUnit.Symbol())
	assert.Equal(t, *newType, newUnit.Type())
	assert.Equal(t, *newSystem, newUnit.System())
}
