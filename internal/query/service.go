// Package query implements all data queries.
package query

// Service is the business logic for queries.
type Service struct {
	repo Repository
}

// NewService creates a new query Service.
func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}
