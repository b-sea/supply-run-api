package query

import "github.com/google/uuid"

type GetAllUnitsHandler struct {
	reader GetAllUnitsReader
}

func (h GetAllUnitsHandler) Handle(userID uuid.UUID) ([]*Unit, error) {
	return h.reader.GetAllUnits(userID)
}

type GetAllUnitsReader interface {
	GetAllUnits(userID uuid.UUID) ([]*Unit, error)
}
