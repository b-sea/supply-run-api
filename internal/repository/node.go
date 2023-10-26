// Package repository defines interfaces for communicating with data layers.
package repository

import "github.com/b-sea/supply-run-api/internal/model"

type NodeRepo[E model.Node, F model.Filter] interface {
	Find(filter *F) ([]*E, error)
	GetOne(id string) (*E, error)
	GetMany(ids []string) ([]*E, error)
	Create(entity E) error
	Update(entity E) error
	Delete(id string) error
}

type IAccountRepo interface {
	NodeRepo[model.Account, model.AccountFilter]
}
