package query

import "github.com/google/uuid"

type GetUnitsHandler struct {
	reader GetUnitsReader
}

func (h GetUnitsHandler) Handle(userID uuid.UUID, unitIDs []uuid.UUID) ([]*Unit, error) {
	return h.reader.GetUnits(userID, unitIDs)
}

type GetUnitsReader interface {
	GetUnits(userID uuid.UUID, unitIDs []uuid.UUID) ([]*Unit, error)
}
