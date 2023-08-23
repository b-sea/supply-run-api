package model

import (
	"time"

	"github.com/google/uuid"
)

const ProductKind Kind = "Product"

type Product struct {
	Metadata

	Name        string   `json:"name"`
	Description string   `json:"description"`
	CategoryID  GlobalID `json:"categoryId"`
}

func (m Product) IsListable()  {}
func (m Product) IsStockable() {}

type CreateProductInput struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	CategoryID  GlobalID `json:"categoryId"`
}

func (m CreateProductInput) ToNode() *Product {
	return &Product{
		Metadata: Metadata{
			ID: GlobalID{
				Key:  uuid.NewString(),
				Kind: ProductKind,
			},
			CreatedBy: "me",
			CreatedAt: time.Now().UTC(),
			UpdatedBy: "me",
			UpdatedAt: time.Now().UTC(),
		},
		Name:        m.Name,
		Description: m.Description,
		CategoryID:  m.CategoryID,
	}
}

func (m CreateProductInput) Validate() error {
	return nil
}

type UpdateProductResult struct {
	ID          GlobalID `json:"id"`
	Revision    string   `json:"revision"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	CategoryID  GlobalID `json:"categoryId"`
}

func (m UpdateProductResult) GetID() GlobalID {
	return m.ID
}

func (m UpdateProductResult) GetRevision() string {
	return m.Revision
}

func (m UpdateProductResult) Validate() error {
	return nil
}

func (m UpdateProductResult) MergeNode(node *Product) {
	node.Revision = m.Revision
	node.UpdatedAt = time.Now().UTC()

	node.Name = m.Name
	node.Description = m.Description
	node.CategoryID = m.CategoryID
}

type ProductFilter struct {
	NodeFilter

	Name        *StringFilter `json:"name"`
	Description *StringFilter `json:"description"`
	CategoryID  *IDFilter     `json:"categoryId"`
}
