// Package main is the startup for the supply run api service.
package main

import (
	"github.com/b-sea/supply-run-api/internal/data/memory"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func main() {
	units := memory.NewUnitRepository()
	logrus.Info(units.GetByOwnerID(uuid.New()))
}
