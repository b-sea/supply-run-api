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

// SignupInput defines all settable properties during account creation.
type SignupInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (m SignupInput) Validate() error {
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

// ToNode converts a SignupInput to an Account node.
func (m SignupInput) ToNode(key string, timestamp time.Time) Account {
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

// MergeNode applies all UpdateAccountInput values to an Account node.
func (m UpdateAccountInput) MergeNode(node *Account, timestamp time.Time) {
	if m.Email != nil {
		node.Email = *m.Email
		node.IsVerified = false
	}

	node.UpdatedAt = timestamp
}

// LoginResult is the response after a successful login.
type LoginResult struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
