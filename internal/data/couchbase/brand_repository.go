package couchbase

import (
	"github.com/b-sea/supply-run-api/internal/model"
	"github.com/couchbase/gocb/v2"
)

const brandType dtoType = "brand"

type brandDTO struct {
	metadata
	shared

	Name string `json:"name"`
}

func (d brandDTO) Metadata() metadata {
	return d.metadata
}

func (d brandDTO) FromNode(node *model.Brand) interface{} {
	return brandDTO{
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
		Name: node.Name,
	}
}

func (d brandDTO) ToNode(metadata metadata) *model.Brand {
	return &model.Brand{
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
		Name: d.Name,
	}
}

type BrandRepository struct {
	*crudRepository[brandDTO, model.Brand]
}

func NewBrandRepository(cluster *gocb.Cluster) *BrandRepository {
	return &BrandRepository{
		crudRepository: &crudRepository[brandDTO, model.Brand]{
			cluster:         cluster,
			scope_name:      entityScopeName,
			collection_name: "brands",
			dtoType:         brandType,
		},
	}
}
