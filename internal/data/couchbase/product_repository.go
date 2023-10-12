package couchbase

import (
	"github.com/b-sea/supply-run-api/internal/model"
	"github.com/couchbase/gocb/v2"
)

const productType dtoType = "product"

type productDTO struct {
	shared

	Name        string `json:"name"`
	Description string `json:"description"`

	CategoryKey string `json:"categoryKey"`
}

func (d productDTO) Metadata() metadata {
	return d.metadata
}

func (d productDTO) FromNode(node *model.Product) interface{} {
	return productDTO{
		shared: shared{
			metadata: metadata{
				Key:      node.ID.Key,
				Revision: stringToCas(node.Revision),
			},
			Type: kindToType(node.ID.Kind),

			CreatedBy: node.CreatedBy,
			CreatedAt: node.CreatedAt,

			UpdatedBy: node.UpdatedBy,
			UpdatedAt: node.UpdatedAt,
		},
		Name:        node.Name,
		Description: node.Description,
		CategoryKey: node.CategoryID.Key,
	}
}

func (d productDTO) ToNode(metadata metadata) *model.Product {
	return &model.Product{
		Metadata: model.Metadata{
			ID: model.GlobalID{
				Key:  metadata.Key,
				Kind: typeToKind(d.Type),
			},
			Revision: casToString(metadata.Revision),

			CreatedBy: d.CreatedBy,
			CreatedAt: d.CreatedAt,

			UpdatedBy: d.UpdatedBy,
			UpdatedAt: d.UpdatedAt,
		},
		Name:        d.Name,
		Description: d.Description,
		CategoryID: model.GlobalID{
			Key:  d.CategoryKey,
			Kind: typeToKind(categoryType),
		},
	}
}

type ProductRepository struct {
	*crudRepository[productDTO, model.Product, model.ProductFilter]
}

func NewProductRepository(cluster *gocb.Cluster) *ProductRepository {
	return &ProductRepository{
		crudRepository: &crudRepository[productDTO, model.Product, model.ProductFilter]{
			cluster:        cluster,
			scopeName:      entityScopeName,
			collectionName: "products",
			dtoType:        productType,
		},
	}
}
