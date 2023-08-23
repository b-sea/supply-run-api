package repository

import (
	"github.com/b-sea/supply-run-api/internal/model"
)

type IReadRepo[E model.Node, F any] interface {
	// Find(filter *F) ([]*E, error)
	GetOne(key string) (*E, error)
	GetMany(keys []string) ([]*E, error)
}

type IReadWriteRepo[E model.Node, F any] interface {
	IReadRepo[E, F]

	Create(node *E) (string, error)
	Update(node *E) (string, error)
	Delete(id string) error
}

type IProductReadRepo interface {
	IReadRepo[model.Product, any]
}

type IProductReadWriteRepo interface {
	IReadWriteRepo[model.Product, any]
}
