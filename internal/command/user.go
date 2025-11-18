package command

import (
	"context"

	"github.com/b-sea/supply-run-api/internal/entity"
	"github.com/b-sea/supply-run-api/internal/user"
)

type CreateUser struct {
	ID       entity.ID
	Username string
}

func (s *Service) CreateUser(ctx context.Context, cmd CreateUser) (entity.ID, error) {
	// TODO: Only admins can do this

	created := user.New(cmd.ID, cmd.Username)

	if err := s.users.CreateUser(ctx, created); err != nil {
		return entity.ID{}, err
	}

	return created.ID(), nil
}

func (s *Service) DeleteUser(ctx context.Context, id entity.ID) error {
	// TODO: Only admins can do this

	if err := s.users.DeleteUser(ctx, id); err != nil {
		return err
	}

	return nil
}
