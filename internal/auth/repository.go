package auth

import (
	"github.com/b-sea/supply-run-api/internal/entity"
)

type Repository interface {
	GetAuthUser(username string) (*User, error)

	GetAccount(id entity.ID) (*Account, error)
	CreateAccount(account *Account) error
	UpdateAccount(account *Account) error
	DeleteAccount(id entity.ID) error
}
