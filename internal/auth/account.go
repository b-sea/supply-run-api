// Package auth defines authentication workflows and entities.
package auth

import (
	"errors"
	"time"

	"github.com/b-sea/supply-run-api/internal/entity"
)

// AccountOption is a creation option for a user Account.
type AccountOption func(account *Account) error

func setPassword(password string) AccountOption {
	return func(account *Account) error {
		account.password = password
		return nil
	}
}

// Account defines exhaustive user information.
type Account struct {
	id        entity.ID
	username  string
	password  string
	createdAt time.Time
	updatedAt time.Time
}

// NewAccount creates a new user Account.
func NewAccount(id entity.ID, username string, password string, timestamp time.Time) *Account {
	return &Account{
		id:        id,
		username:  username,
		password:  password,
		createdAt: timestamp,
		updatedAt: timestamp,
	}
}

// Update an existing user Account.
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

// ID returns the Account ID.
func (a *Account) ID() entity.ID {
	return a.id
}

// Username returns the Account username.
func (a *Account) Username() string {
	return a.username
}

// Password returns the Account hashed password.
func (a *Account) Password() string {
	return a.password
}

// CreatedAt returns when the Account was created.
func (a *Account) CreatedAt() time.Time {
	return a.createdAt
}

// UpdatedAt returns the Account was updated.
func (a *Account) UpdatedAt() time.Time {
	return a.updatedAt
}
