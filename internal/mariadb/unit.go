package mariadb

import (
	"context"

	"github.com/b-sea/supply-run-api/internal/entity"
	"github.com/b-sea/supply-run-api/internal/query"
)

func (r *Repository) GetUnits(ctx context.Context, ids []entity.ID) ([]*query.Unit, error) {
	return nil, nil
}
