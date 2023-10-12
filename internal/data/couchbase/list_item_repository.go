package couchbase

import (
	"github.com/b-sea/supply-run-api/internal/model"
	"github.com/couchbase/gocb/v2"
)

const listItemType dtoType = "list_item"

type listItemDTO struct {
	shared

	Note string `json:"note"`

	IsImportant bool `json:"isImportant"`
	IsComplete  bool `json:"isComplete"`

	ShoppingListKey string `json:"shoppingListKey"`
	Item            item   `json:"item"`
}

type item struct {
	Key  string  `json:"key"`
	Type dtoType `json:"type"`
}

func (d listItemDTO) FromNode(node *model.ListItem) interface{} {
	return listItemDTO{
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
		Note: node.Note,

		IsImportant: node.IsImportant,
		IsComplete:  node.IsComplete,

		ShoppingListKey: node.ShoppingListID.Key,
		Item: item{
			Key:  node.ItemID.Key,
			Type: kindToType(node.ItemID.Kind),
		},
	}
}

func (d listItemDTO) ToNode(metadata metadata) *model.ListItem {
	return &model.ListItem{
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
		Note: d.Note,

		IsImportant: d.IsImportant,
		IsComplete:  d.IsComplete,

		ShoppingListID: model.GlobalID{
			Key:  d.ShoppingListKey,
			Kind: typeToKind(shoppingListType),
		},
		ItemID: model.GlobalID{
			Key:  d.Item.Key,
			Kind: typeToKind(d.Item.Type),
		},
	}
}

type ListItemRepository struct {
	*crudRepository[listItemDTO, model.ListItem, any]
}

func NewListItemRepository(cluster *gocb.Cluster) *ListItemRepository {
	return &ListItemRepository{
		crudRepository: &crudRepository[listItemDTO, model.ListItem, any]{
			cluster:        cluster,
			scopeName:      functionScopeName,
			collectionName: "list_items",
			dtoType:        listItemType,
		},
	}
}
