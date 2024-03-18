package unit_test

import (
	"testing"

	"github.com/b-sea/supply-run-api/internal/domain/unit"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func TestUnitCreate(t *testing.T) {
	t.Parallel()

	admin := uuid.New()

	imperial, err := unit.NewSystem(
		unit.NewSystemInput{
			Owner: admin,
			Name:  "imperial",
		},
	)
	if err != nil {
		panic(err)
	}
	logrus.Infof("%+v", imperial)

	metric, err := unit.NewSystem(
		unit.NewSystemInput{
			Owner: admin,
			Name:  "metric",
		},
	)
	if err != nil {
		panic(err)
	}
	logrus.Infof("%+v", metric)

	volume, err := unit.NewType(
		unit.NewTypeInput{
			Owner: admin,
			Name:  "volume",
		},
	)
	if err != nil {
		panic(err)
	}
	logrus.Infof("%+v", volume)

	mass, err := unit.NewType(
		unit.NewTypeInput{
			Owner: admin,
			Name:  "mass",
		},
	)
	if err != nil {
		panic(err)
	}
	logrus.Infof("%+v", mass)

	ml, err := unit.NewUnit(
		unit.NewUnitInput{
			Owner:  admin,
			Name:   "millilitre",
			Symbol: "ml",
			System: *metric,
			Type:   *volume,
		},
	)
	if err != nil {
		panic(err)
	}
	logrus.Infof("%+v", ml)

	tsp, err := unit.NewUnit(
		unit.NewUnitInput{
			Owner:  admin,
			Name:   "teaspoon",
			Symbol: "tsp",
			System: *imperial,
			Type:   *volume,
			ConvertFrom: &unit.ConversionInput{
				Unit:  *ml,
				Ratio: 4.929,
			},
		},
	)
	if err != nil {
		panic(err)
	}
	logrus.Infof("%+v", tsp)

	tbsp, err := unit.NewUnit(
		unit.NewUnitInput{
			Owner:  admin,
			Name:   "tablespoon",
			Symbol: "tbsp",
			System: *imperial,
			Type:   *volume,
			ConvertFrom: &unit.ConversionInput{
				Unit:  *tsp,
				Ratio: 2,
			},
		},
	)
	if err != nil {
		panic(err)
	}
	logrus.Infof("%+v", tbsp)
}
