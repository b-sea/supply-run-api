package model

import (
	"time"
)

type Kind string

const (
	ShoppingListKind Kind = "ShoppingList"
	ListItemKind     Kind = "ListItem"
	CategoryKind     Kind = "Category"
	BrandProductKind Kind = "BrandProduct"
	BrandKind        Kind = "Brand"
)

type Node interface {
	GetID() GlobalID
	GetRevision() string
	IsNode()
}

type GlobalID struct {
	Key  string
	Kind Kind
}

type Metadata struct {
	ID       GlobalID `json:"id"`
	Revision string   `json:"revision"`

	CreatedBy string    `json:"createdBy"`
	CreatedAt time.Time `json:"createdAt"`

	UpdatedBy string    `json:"updatedBy"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (m Metadata) IsNode()             {}
func (m Metadata) GetID() GlobalID     { return m.ID }
func (m Metadata) GetRevision() string { return m.Revision }

type CreateInput[N Node] interface {
	Validate() error
	ToNode() *N
}

type CreateResult struct {
	ID       GlobalID `json:"id"`
	Revision string   `json:"revision"`
}

type UpdateInput[N Node] interface {
	GetID() GlobalID
	GetRevision() string
	Validate() error
	MergeNode(node *N)
}

type UpdateResult struct {
	ID       GlobalID `json:"id"`
	Revision string   `json:"revision"`
}

type DeleteInput struct {
	ID       GlobalID `json:"id"`
	Revision string   `json:"revision"`
}

type ShoppingList struct {
	Metadata

	Name        string `json:"name"`
	Description string `json:"description"`
}

type Listable interface {
	IsListable()
}

type Stockable interface {
	IsStockable()
}

type ListItem struct {
	Metadata

	IsImportant bool   `json:"isImportant"`
	IsComplete  bool   `json:"isComplete"`
	Note        string `json:"name"`

	ShoppingListID GlobalID `json:"shoppingListId"`
	ItemID         GlobalID `json:"itemId"`
}

type Category struct {
	Metadata

	Name string `json:"name"`
}

type Brand struct {
	Metadata

	Name string `json:"name"`
}

type BrandProduct struct {
	Metadata

	Name        string `json:"name"`
	Description string `json:"description"`

	BrandID   GlobalID `json:"brandId"`
	ProductID GlobalID `json:"productId"`
}

func (m BrandProduct) IsListable() bool { return true }
