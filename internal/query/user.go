package query

import (
	"context"

	"github.com/b-sea/supply-run-api/internal/entity"
)

// GetUser returns a single user from an id.
func (s *Service) GetUser(ctx context.Context, id entity.ID) (*User, error) {
	found, err := s.GetUsers(ctx, []entity.ID{id})
	if err != nil {
		return nil, err
	}

	if len(found) == 0 {
		return nil, entity.ErrNotFound
	}

	return found[0], nil
}

// GetUsers returns multiple users from a list of ids.
func (s *Service) GetUsers(ctx context.Context, ids []entity.ID) ([]*User, error) {
	found, err := s.repo.GetUsers(ctx, ids)
	if err != nil {
		return nil, queryError(err)
	}

	return found, nil
}
