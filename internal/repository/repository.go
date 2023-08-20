package repository

import (
	"github.com/b-sea/supply-run-api/internal/model"
)

type ICRUDRepository[D any, E any, F any] interface {
	Find(filter *F) ([]E, error)
	GetOne(key string) (E, error)
	GetMany(keys []string) ([]E, error)

	Create(entity E) (string, error)
	Update(entity E) (string, error)
	Delete(id any) error
}

type IProductRepository[D any, E model.Product, F model.NodeFilter] interface {
	ICRUDRepository[D, E, F]
}

type IStoreRepository[D any, E model.Store, F model.NodeFilter] interface {
	ICRUDRepository[D, E, F]
}

type ILocationRepository[D any, E model.Location, F model.NodeFilter] interface {
	ICRUDRepository[D, E, F]
}

type IShoppingListRepository[D any, E model.ShoppingList, F model.NodeFilter] interface {
	ICRUDRepository[D, E, F]
}

type IListItemRepository[D any, E model.ListItem, F model.NodeFilter] interface {
	ICRUDRepository[D, E, F]
}
