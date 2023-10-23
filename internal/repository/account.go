package repository

import "github.com/b-sea/supply-run-api/internal/model"

type IAccountRepo interface {
	Find(filter *model.AccountFilter) ([]*model.Account, error)
	GetOne(id string) (*model.Account, error)
	GetMany(ids []string) ([]*model.Account, error)
	Create(entity *model.Account) (string, error)
	Update(entity *model.Account) (string, error)
	Delete(id string) error
}
