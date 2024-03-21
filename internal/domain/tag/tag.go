// Package tag defines everything to manage the item tagging.
package tag

import (
	"time"

	"github.com/b-sea/supply-run-api/internal/domain"
	"github.com/google/uuid"
)

// Tag is categorizing data.
type Tag struct {
	id        uuid.UUID
	ownerID   uuid.UUID
	createdAt time.Time
	name      string
}

// Name returns the tag name.
func (t *Tag) Name() string {
	return t.name
}

// OwnerID returns the creator of the tag.
func (t *Tag) OwnerID() uuid.UUID {
	return t.ownerID
}

// NewTag creates a new tag.
func NewTag(name string, ownerID uuid.UUID) (*Tag, error) {
	if name == "" {
		return nil, &domain.ValidationError{Issues: []string{"tag cannot be empty"}}
	}

	return &Tag{
		id:        uuid.New(),
		ownerID:   ownerID,
		createdAt: time.Now().UTC(),
		name:      name,
	}, nil
}

// Hydrate returns a tag in an existing state.
func Hydrate(id uuid.UUID, name string, createdAt time.Time, ownerID uuid.UUID) (*Tag, error) {
	tag, err := NewTag(name, ownerID)
	if err != nil {
		return nil, err
	}

	tag.id = id
	tag.createdAt = createdAt

	return tag, nil
}
