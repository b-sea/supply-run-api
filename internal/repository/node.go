// Package repository defines interfaces for communicating with data layers.
package repository

import "github.com/b-sea/supply-run-api/internal/model"

type INode[N model.Node, F model.Filter] interface {
	Find(filter *F, accountID string) ([]*N, error)
	GetByIDs(ids []string, accountID string) ([]*N, error)
	Create(node N, accountID string) error
	Update(node N, accountID string) error
	Delete(id string, accountID string) error
}

type IAccount interface {
	GetByID(id string) (*model.Account, error)
	GetByEmail(email string) (*model.Account, error)
	Create(account model.Account) error
	Update(account model.Account) error
	Delete(id string) error
}

type IRecipe interface {
	INode[model.Recipe, model.RecipeFilter]
}
