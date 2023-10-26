// Package service implements all business logic for the API.
package service

import "github.com/b-sea/supply-run-api/internal/model"

type INodeService[N model.Node, F model.Filter, C model.CreateInput[N], U model.UpdateInput[N]] interface {
	Find(filter *F) ([]*N, error)
	GetOne(id model.ID) (*N, error)
	GetMany(id []model.ID) ([]*N, error)
	Create(input C) error
	Update(input U) error
	Delete(id model.ID) error
}
