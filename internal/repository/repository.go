package repository

import (
	"github.com/b-sea/supply-run-api/internal/model"
)

type IReadRepo[N model.Node, F any] interface {
	Find(filter F) ([]*N, error)
	GetOne(key string) (*N, error)
	GetMany(keys []string) ([]*N, error)
}

type IReadWriteRepo[N model.Node, F any] interface {
	IReadRepo[N, F]

	Create(node *N) (string, error)
	Update(node *N) (string, error)
	Delete(id string) error
}

type IProductReadRepo interface {
	IReadRepo[model.Product, model.ProductFilter]
}

type IProductReadWriteRepo interface {
	IReadWriteRepo[model.Product, model.ProductFilter]
}
