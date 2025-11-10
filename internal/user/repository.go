package user

import (
	"context"

	"github.com/b-sea/supply-run-api/internal/entity"
)

// Repository defines all data interactions required for users.
type Repository interface {
	CreateUser(ctx context.Context, user *User) error
	DeleteUser(ctx context.Context, id entity.ID) error
}
