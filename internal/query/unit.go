package query

import (
	"context"

	"github.com/b-sea/supply-run-api/internal/entity"
)

// GetUnits returns multiple units from a list of ids.
func (s *Service) GetUnits(ctx context.Context, ids []entity.ID) ([]*Unit, error) {
	found, err := s.units.GetUnits(ctx, ids)
	if err != nil {
		return nil, queryError(err)
	}

	return found, nil
}
