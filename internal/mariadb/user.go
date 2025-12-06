package mariadb

import (
	"context"

	"github.com/b-sea/supply-run-api/internal/entity"
	"github.com/b-sea/supply-run-api/internal/query"
)

func (r *Repository) GetUsers(ctx context.Context, ids []entity.ID) ([]*query.User, error) {
	return nil, nil
}
