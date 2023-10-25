// Package model defines all data entities shared between front end, service, and repository layers.
package model

import "time"

// AccountKind is the ID type associated with Accounts.
const AccountKind = kind("Account")

// Account defines all account properties.
type Account struct {
	Entity

	Username string `json:"username"`
	Password string `json:"-"`
	Email    string `json:"email"`

	IsVerified bool       `json:"isVerified"`
	LastLogin  *time.Time `json:"lastLogin"`
}

// AccountFilter defines all searchable account properties.
type AccountFilter struct {
	Username *StringFilter `json:"username"`
}

// CreateAccountInput defines all settable properties during account creation.
type CreateAccountInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// ToEntity converts a CreateAccountInput to an Account entity.
func (m CreateAccountInput) ToEntity(key string, timestamp time.Time) Account {
	return Account{
		Entity: Entity{
			ID: ID{
				Kind: AccountKind,
				Key:  key,
			},
			CreatedAt: timestamp,
			UpdatedAt: timestamp,
		},
		Username: m.Username,
		Password: m.Password,
		Email:    m.Email,
	}
}

// UpdateAccountInput defines all settable properties during account update.
type UpdateAccountInput struct {
	Email    *string `json:"email"`
	Password *string `json:"password"`
}

// MergeEntity applies all UpdateAccountInput values to an Account entity.
func (m UpdateAccountInput) MergeEntity(entity *Account, timestamp time.Time) {
	if m.Email != nil {
		entity.Email = *m.Email
	}

	entity.UpdatedAt = timestamp
}

// LoginResult is the response after a successful login.
type LoginResult struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
