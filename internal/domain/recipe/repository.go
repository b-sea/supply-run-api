// Package recipe defines everything to manage the recipes domain.
package recipe

import "github.com/google/uuid"

// Repository defines all functions required to interact with recipes.
type Repository interface {
	GetByOwnerIDs(ids []uuid.UUID) ([]*Recipe, error)
	Create(recipe *Recipe) error
	Update(recipe *Recipe) error
	Delete(id uuid.UUID) error
}
