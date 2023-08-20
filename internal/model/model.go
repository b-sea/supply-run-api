package model

import (
	"time"
)

type Entity interface {
	IsNode() bool
}

type GlobalID struct {
	Key  string
	Kind string
}

type Metadata struct {
	ID       GlobalID `json:"id"`
	Revision string   `json:"revision"`

	CreatedBy string    `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`

	UpdatedBy string    `json:"updated_by"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (m Metadata) IsNode() bool { return true }

type Product struct {
	Metadata

	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
}

func (m Product) IsNode() bool { return true }

type Store struct {
	Metadata

	Name    string `json:"name"`
	Address string `json:"address"`
	Website string `json:"website"`
}

type Location struct {
	Metadata

	Name      string   `json:"name"`
	ProductID GlobalID `json:"product_id"`
	StoreID   GlobalID `json:"store_id"`
}

type ShoppingList struct {
	Metadata

	Name    string   `json:"name"`
	StoreID GlobalID `json:"store_id"`
}

type ListItem struct {
	Metadata

	Name        string `json:"name"`
	IsImportant bool   `json:"is_important"`
	IsComplete  bool   `json:"is_complete"`

	ShoppingListID GlobalID `json:"shopping_list_id"`
	ProductID      GlobalID `json:"product_id"`
}

type NodeFilter struct {
	ID   *IDFilter     `json:"id"`
	Name *StringFilter `json:"name"`

	CreatedBy *StringFilter `json:"created_by"`
	CreatedAt *TimeFilter   `json:"created_at"`

	UpdatedBy *StringFilter `json:"updated_by"`
	UpdatedAt *TimeFilter   `json:"updated_at"`
}
