// Package cookbook defines everything to manage recipes.
package cookbook

import (
	"fmt"

	"github.com/b-sea/supply-run-api/internal/domain/recipe"
	"github.com/b-sea/supply-run-api/internal/domain/unit"
	"github.com/google/uuid"
)

// Config is a unit service configuration option.
type Config func(*Service) error

// Service is a cookbook service.
type Service struct {
	recipes recipe.Repository
	units   unit.Repository
}

// NewService creates a new cookbook service.
func NewService(recipes recipe.Repository, units unit.Repository) *Service {
	return &Service{
		recipes: recipes,
		units:   units,
	}
}

// GetUnits returns all units owned by the given users.
func (s *Service) GetUnits(owners []uuid.UUID) ([]*unit.Unit, error) {
	results, err := s.units.GetByOwnerIDs(owners)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return results, nil
}
