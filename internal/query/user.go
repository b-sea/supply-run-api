package query

import (
	"context"

	"github.com/b-sea/supply-run-api/internal/entity"
)

// GetUsers returns multiple users from a list of ids.
func (s *Service) GetUsers(ctx context.Context, ids []entity.ID) ([]*User, error) {
	found, err := s.users.GetUsers(ctx, ids)
	if err != nil {
		return nil, queryError(err)
	}

	return found, nil
}
