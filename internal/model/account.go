// Package model defines all data entities shared between front end, service, and repository layers.
package model

import (
	"net/mail"
	"time"
)

// AccountKind is the ID type associated with Accounts.
const AccountKind = kind("Account")

// Account defines all account properties.
type Account struct {
	ID        ID        `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	Email    string `json:"email"`
	Password string `json:"-"`

	IsVerified bool       `json:"isVerified"`
	LastLogin  *time.Time `json:"lastLogin"`
}

// GetID returns the ID of an Account.
func (m Account) GetID() ID {
	return m.ID
}

// AccountFilter defines all searchable account properties.
type AccountFilter struct {
	Email *StringFilter `json:"email"`
}

func (m AccountFilter) IsFilter() {}

// CreateAccountInput defines all settable properties during account creation.
type CreateAccountInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (m CreateAccountInput) Validate() error {
	issues := make([]string, 0)
	if _, err := mail.ParseAddress(m.Email); err != nil {
		issues = append(issues, "invalid email format")
	}

	if len(issues) > 0 {
		return ValidationError{
			Issues: issues,
		}
	}
	return nil
}

// ToEntity converts a CreateAccountInput to an Account entity.
func (m CreateAccountInput) ToEntity(key string, timestamp time.Time) Account {
	return Account{
		ID: ID{
			Kind: AccountKind,
			Key:  key,
		},
		CreatedAt: timestamp,
		UpdatedAt: timestamp,

		Password: m.Password,
		Email:    m.Email,
	}
}

// UpdateAccountInput defines all settable properties during account update.
type UpdateAccountInput struct {
	Email    *string `json:"email"`
	Password *string `json:"password"`
}

func (m UpdateAccountInput) Validate() error {
	return nil
}

// MergeEntity applies all UpdateAccountInput values to an Account entity.
func (m UpdateAccountInput) MergeEntity(entity *Account, timestamp time.Time) {
	if m.Email != nil {
		entity.Email = *m.Email
		entity.IsVerified = false
	}

	entity.UpdatedAt = timestamp
}

// LoginResult is the response after a successful login.
type LoginResult struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
