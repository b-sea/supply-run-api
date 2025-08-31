package mock

import (
	"github.com/b-sea/supply-run-api/internal/auth"
	"github.com/b-sea/supply-run-api/internal/entity"
)

var _ auth.Repository = (*Auth)(nil)

type Auth struct {
	GetAuthUserResult *auth.User
	GetAuthUserErr    error

	GetAccountResult *auth.Account
	GetAccountErr    error
	CreateAccountErr error
	UpdateAccountErr error
	DeleteAccountErr error
}

func (a *Auth) GetAuthUser(username string) (*auth.User, error) {
	return a.GetAuthUserResult, a.GetAuthUserErr
}

func (a *Auth) GetAccount(id entity.ID) (*auth.Account, error) {
	return a.GetAccountResult, a.GetAccountErr
}

func (a *Auth) CreateAccount(account *auth.Account) error {
	return a.CreateAccountErr
}

func (a *Auth) UpdateAccount(account *auth.Account) error {
	return a.UpdateAccountErr
}

func (a *Auth) DeleteAccount(id entity.ID) error {
	return a.DeleteAccountErr
}
