// Package recipe defines everything to manage the recipes domain.
package recipe

import (
	"github.com/google/uuid"
)

type Snippet struct {
	id      uuid.UUID
	ownerID uuid.UUID
	name    string
	desc    string
	tagIDs  []uuid.UUID
}

func (s *Snippet) ID() uuid.UUID {
	return s.id
}

func (s *Snippet) OwnerID() uuid.UUID {
	return s.ownerID
}

func (s *Snippet) Name() string {
	return s.name
}

func (s *Snippet) Description() string {
	return s.desc
}

func (s *Snippet) TagIDs() []uuid.UUID {
	return s.tagIDs
}
