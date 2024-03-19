// Package recipe defines everything to manage the recipes domain.
package recipe

import "github.com/google/uuid"

// Repository defines all functions required to interact with recipes.
type Repository interface {
	Find(filter *Filter) ([]*Recipe, error)
	GetOne(id uuid.UUID) (*Recipe, error)
	Create(recipe *Recipe) error
	Update(recipe *Recipe) error
	Delete(id uuid.UUID) error
}

// Filter is a search filter for recipes.
type Filter struct {
	Owners []uuid.UUID
	Name   *string
	Tags   []*Tag
}
