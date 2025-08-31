// Package memory implements an in-memory data store.
package memory

import (
	"github.com/b-sea/supply-run-api/internal/auth"
	"github.com/b-sea/supply-run-api/internal/entity"
)

var _ auth.Repository = (*Repository)(nil)

// Option is a creation option for a Repository.
type Option func(repo *Repository)

// WithAccounts pre-defines user Accounts that will exist in the data store.
func WithAccounts(accounts ...*auth.Account) Option {
	return func(repo *Repository) {
		for i := range accounts {
			_ = repo.CreateAccount(accounts[i])
		}
	}
}

// Repository is an in-memory data store.
type Repository struct {
	usernameIDs map[string]entity.ID
	accounts    map[entity.ID]*auth.Account
}

// NewRepository creates a new Repository.
func NewRepository(options ...Option) *Repository {
	repo := &Repository{
		usernameIDs: make(map[string]entity.ID),
		accounts:    make(map[entity.ID]*auth.Account),
	}

	for _, option := range options {
		option(repo)
	}

	return repo
}

// GetAuthUser returns a minimal User from a username.
func (r *Repository) GetAuthUser(username string) (*auth.User, error) {
	id, ok := r.usernameIDs[username]
	if !ok {
		return nil, auth.ErrNotFound
	}

	found, ok := r.accounts[id]
	if !ok {
		return nil, auth.ErrNotFound
	}

	return auth.NewUser(found.ID(), found.Username()), nil
}

// GetAccount returns a user Account from an ID.
func (r *Repository) GetAccount(id entity.ID) (*auth.Account, error) {
	found, ok := r.accounts[id]
	if !ok {
		return nil, auth.ErrNotFound
	}

	return found, nil
}

// CreateAccount creates a new user Account.
func (r *Repository) CreateAccount(account *auth.Account) error {
	if _, ok := r.accounts[account.ID()]; ok {
		return auth.ErrDuplicateAccount
	}

	r.usernameIDs[account.Username()] = account.ID()
	r.accounts[account.ID()] = account

	return nil
}

// UpdateAccount updates an existing user Account.
func (r *Repository) UpdateAccount(account *auth.Account) error {
	if _, ok := r.accounts[account.ID()]; !ok {
		return auth.ErrNotFound
	}

	r.usernameIDs[account.Username()] = account.ID()
	r.accounts[account.ID()] = account

	return nil
}

// DeleteAccount deletes an existing user Account.
func (r *Repository) DeleteAccount(id entity.ID) error {
	found, ok := r.accounts[id]
	if !ok {
		return nil
	}

	delete(r.usernameIDs, found.Username())
	delete(r.accounts, id)

	return nil
}
