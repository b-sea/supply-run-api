package auth

import (
	"errors"
	"time"

	"github.com/b-sea/supply-run-api/internal/entity"
)

type AccountOption func(account *Account) error

func setPassword(password string) AccountOption {
	return func(account *Account) error {
		account.password = password
		return nil
	}
}

type Account struct {
	id        entity.ID
	username  string
	password  string
	createdAt time.Time
	updatedAt time.Time
}

func NewAccount(id entity.ID, username string, password string, timestamp time.Time) *Account {
	return &Account{
		id:        id,
		username:  username,
		password:  password,
		createdAt: timestamp,
		updatedAt: timestamp,
	}
}

func (a *Account) Update(timestamp time.Time, options ...AccountOption) error {
	errs := make([]error, 0)

	for _, option := range options {
		if err := option(a); err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	a.updatedAt = timestamp

	return nil
}

func (a *Account) ID() entity.ID {
	return a.id
}

func (a *Account) Username() string {
	return a.username
}

func (a *Account) Password() string {
	return a.password
}

func (a *Account) CreatedAt() time.Time {
	return a.createdAt
}

func (a *Account) UpdatedAt() time.Time {
	return a.updatedAt
}
