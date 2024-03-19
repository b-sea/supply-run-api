// Package cookbook defines everything to manage recipes.
package cookbook

import (
	"fmt"

	"github.com/b-sea/supply-run-api/internal/domain/recipe"
	"github.com/b-sea/supply-run-api/internal/domain/unit"
	"github.com/google/uuid"
)

// Option is a cookbook service configuration option.
type Option func(*Service)

// WithBaseUser declares a shared user on the service to access default data.
func WithBaseUser(id uuid.UUID) Option {
	return func(s *Service) {
		s.baseUser = id
	}
}

// Service is a cookbook service.
type Service struct {
	recipes  recipe.Repository
	units    unit.Repository
	baseUser uuid.UUID
}

// NewService creates a new cookbook service.
func NewService(recipes recipe.Repository, units unit.Repository, opts ...Option) *Service {
	service := &Service{
		baseUser: uuid.Nil,
		recipes:  recipes,
		units:    units,
	}

	for _, opt := range opts {
		opt(service)
	}

	return service
}

// FindRecipes searches for recipes.
func (s *Service) FindRecipes(filter *recipe.Filter) ([]*recipe.Recipe, error) {
	if filter == nil {
		filter = &recipe.Filter{}
	}

	if filter.Owners == nil {
		filter.Owners = make([]uuid.UUID, 0)
	}

	if s.baseUser != uuid.Nil {
		filter.Owners = append(filter.Owners, s.baseUser)
	}

	results, err := s.recipes.Find(filter)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return results, nil
}

// FindUnits searches for units.
func (s *Service) FindUnits(filter *unit.Filter) ([]*unit.Unit, error) {
	if filter == nil {
		filter = &unit.Filter{}
	}

	if filter.Owners == nil {
		filter.Owners = make([]uuid.UUID, 0)
	}

	if s.baseUser != uuid.Nil {
		filter.Owners = append(filter.Owners, s.baseUser)
	}

	results, err := s.units.Find(filter)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return results, nil
}
