package couchbase

// import (
// 	"github.com/b-sea/supply-run-api/internal/model"
// 	"github.com/couchbase/gocb/v2"
// )

const shoppingListType dtoType = "shopping_list"

// type shoppingListDTO struct {
// 	metadata
// 	shared

// 	Name        string `json:"name"`
// 	Description string `json:"description"`
// }

// func (d shoppingListDTO) FromNode(node *model.ShoppingList) interface{} {
// 	return shoppingListDTO{
// 		metadata: metadata{
// 			Key:      node.ID.Key,
// 			Revision: stringToCas(node.Revision),
// 		},
// 		shared: shared{
// 			Type: kindToType(node.ID.Kind),

// 			CreatedBy: node.CreatedBy,
// 			CreatedAt: node.CreatedAt,

// 			UpdatedBy: node.UpdatedBy,
// 			UpdatedAt: node.UpdatedAt,
// 		},

// 		Name:        node.Name,
// 		Description: node.Description,
// 	}
// }

// func (d shoppingListDTO) ToNode(metadata metadata) *model.ShoppingList {
// 	return &model.ShoppingList{
// 		Metadata: model.Metadata{
// 			ID: model.GlobalID{
// 				Key:  metadata.Key,
// 				Kind: typeToKind(d.Type),
// 			},
// 			Revision: casToString(metadata.Revision),

// 			CreatedBy: d.CreatedBy,
// 			CreatedAt: d.CreatedAt,

// 			UpdatedBy: d.UpdatedBy,
// 			UpdatedAt: d.UpdatedAt,
// 		},
// 		Name:        d.Name,
// 		Description: d.Description,
// 	}
// }

// type ShoppingListRepository struct {
// 	*crudRepository[shoppingListDTO, model.ShoppingList]
// }

// func NewShoppingListRepository(cluster *gocb.Cluster) *ShoppingListRepository {
// 	return &ShoppingListRepository{
// 		crudRepository: &crudRepository[shoppingListDTO, model.ShoppingList]{
// 			cluster:         cluster,
// 			scopeName:      functionScopeName,
// 			collectionName: "shopping_lists",
// 			dtoType:         shoppingListType,
// 		},
// 	}
// }
