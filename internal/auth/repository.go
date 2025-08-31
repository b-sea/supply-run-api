// Package auth defines authentication workflows and entities.
package auth

import (
	"github.com/b-sea/supply-run-api/internal/entity"
)

// Repository defines all functions required for authentication.
type Repository interface {
	GetAuthUser(username string) (*User, error)

	GetAccount(id entity.ID) (*Account, error)
	CreateAccount(account *Account) error
	UpdateAccount(account *Account) error
	DeleteAccount(id entity.ID) error
}
