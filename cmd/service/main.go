// Package main is the startup for the supply run api service.
package main

import (
	"github.com/b-sea/supply-run-api/configs"
	"github.com/b-sea/supply-run-api/internal/data/memory"
	"github.com/b-sea/supply-run-api/internal/domain/recipe"
	"github.com/b-sea/supply-run-api/internal/domain/unit"
	"github.com/b-sea/supply-run-api/internal/service/cookbook"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func main() {
	owner := uuid.New()
	units := cookbook.NewService(
		memory.NewRecipeRepository([]*recipe.Recipe{}),
		memory.NewUnitRepository(
			unitCatalog(owner, configs.LoadDefaults()),
		),
	)

	foundUnits, err := units.GetUnits([]uuid.UUID{owner})
	if err != nil {
		logrus.Fatal(err)
	}

	for _, u := range foundUnits {
		logrus.Infof("%s: %s (%s) %s %s", u.ID(), u.Name, u.Symbol, u.System, u.Type)
	}
}

func unitCatalog(owner uuid.UUID, defaults *configs.Defaults) []*unit.Unit {
	results := []*unit.Unit{}

	for _, u := range defaults.Units { //nolint: varnamelen
		results = append(
			results,
			unit.NewUnit(
				u.Name,
				owner,
				unit.WithID(u.ID),
				unit.WithSymbol(u.Symbol),
				unit.WithSystem(unit.SystemFromString(u.System)),
				unit.WithType(unit.TypeFromString(u.Type)),
			),
		)
	}

	return results
}
