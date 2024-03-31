// Package recipe defines everything to manage the recipes domain.
package recipe

import "github.com/google/uuid"

type Repository interface {
	Create(entity *Recipe) error
	Update(entity *Recipe) error
	Delete(id uuid.UUID) error
}
