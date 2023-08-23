package service

import (
	"github.com/b-sea/supply-run-api/internal/model"
	"github.com/b-sea/supply-run-api/internal/repository"
)

type ProductService struct {
	crudService[model.Product, model.ProductFilter, model.CreateProductInput, model.UpdateProductResult]
}

func NewProductService(repo repository.IProductReadWriteRepo) *ProductService {
	return &ProductService{
		crudService[model.Product, model.ProductFilter, model.CreateProductInput, model.UpdateProductResult]{
			repo: repo,
			kind: model.ProductKind,
		},
	}
}
