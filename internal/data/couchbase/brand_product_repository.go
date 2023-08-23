package couchbase

import (
	"github.com/b-sea/supply-run-api/internal/model"
	"github.com/couchbase/gocb/v2"
)

const brandProductType dtoType = "brand_product"

type brandProductDTO struct {
	metadata
	shared

	Name        string `json:"name"`
	Description string `json:"description"`

	BrandKey   string `json:"brandKey"`
	ProductKey string `json:"productKey"`
}

func (d brandProductDTO) Metadata() metadata {
	return d.metadata
}

func (d brandProductDTO) FromNode(node *model.BrandProduct) interface{} {
	return brandProductDTO{
		metadata: metadata{
			Key:      node.ID.Key,
			Revision: stringToCas(node.Revision),
		},
		shared: shared{
			Type: kindToType(node.ID.Kind),

			CreatedBy: node.CreatedBy,
			CreatedAt: node.CreatedAt,

			UpdatedBy: node.UpdatedBy,
			UpdatedAt: node.UpdatedAt,
		},
		Name:        node.Name,
		Description: node.Description,
		BrandKey:    node.BrandID.Key,
		ProductKey:  node.ProductID.Key,
	}
}

func (d brandProductDTO) ToNode(metadata metadata) *model.BrandProduct {
	return &model.BrandProduct{
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

		BrandID: model.GlobalID{
			Key:  d.BrandKey,
			Kind: typeToKind(brandType),
		},
		ProductID: model.GlobalID{
			Key:  d.ProductKey,
			Kind: typeToKind(productType),
		},
	}
}

type BrandProductRepository struct {
	*crudRepository[brandProductDTO, model.BrandProduct]
}

func NewBrandProductRepository(cluster *gocb.Cluster) *BrandProductRepository {
	return &BrandProductRepository{
		crudRepository: &crudRepository[brandProductDTO, model.BrandProduct]{
			cluster:         cluster,
			scope_name:      entityScopeName,
			collection_name: "brand_products",
			dtoType:         brandProductType,
		},
	}
}
