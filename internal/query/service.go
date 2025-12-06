// Package query implements all data queries.
package query

// Service is the business logic for queries.
type Service struct {
	recipes RecipeRepository
	units   UnitRepository
	users   UserRepository
}

// NewService creates a new query Service.
func NewService(recipes RecipeRepository, units UnitRepository, users UserRepository) *Service {
	return &Service{
		recipes: recipes,
		units:   units,
		users:   users,
	}
}
