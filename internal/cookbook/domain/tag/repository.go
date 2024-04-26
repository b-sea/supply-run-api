// Package tag defines everything to manage the item tagging.
package tag

import "github.com/google/uuid"

type Repository interface {
	Create(tag *Tag) error
	Delete(id uuid.UUID) error
}
