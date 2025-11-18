package command

import (
	"time"

	"github.com/b-sea/supply-run-api/internal/entity"
	"github.com/b-sea/supply-run-api/internal/recipe"
	"github.com/b-sea/supply-run-api/internal/user"
)

type Service struct {
	recipes     recipe.Repository
	users       user.Repository
	idFn        func() entity.ID
	timestampFn func() time.Time
}

func NewService(recipes recipe.Repository, users user.Repository) *Service {
	return &Service{
		recipes:     recipes,
		users:       users,
		idFn:        entity.NewRandomID,
		timestampFn: time.Now,
	}
}
