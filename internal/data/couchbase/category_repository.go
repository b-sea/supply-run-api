package couchbase

import (
	"github.com/b-sea/supply-run-api/internal/model"
	"github.com/couchbase/gocb/v2"
)

const categoryType dtoType = "category"

type categoryDTO struct {
	metadata
	shared

	Name string `json:"name"`
}

func (d categoryDTO) Metadata() metadata {
	return d.metadata
}

func (d categoryDTO) FromNode(node *model.Category) interface{} {
	return categoryDTO{
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

func (d categoryDTO) ToNode(metadata metadata) *model.Category {
	return &model.Category{
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

type CategoryRepository struct {
	*crudRepository[categoryDTO, model.Category, any]
}

func NewCategoryRepository(cluster *gocb.Cluster) *CategoryRepository {
	return &CategoryRepository{
		crudRepository: &crudRepository[categoryDTO, model.Category, any]{
			cluster:        cluster,
			scopeName:      entityScopeName,
			collectionName: "categories",
			dtoType:        categoryType,
		},
	}
}
