// Package auth implements authentication workflows and entities.
package auth

import "github.com/b-sea/supply-run-api/internal/entity"

// Tokens defines authentication tokens.
type Tokens struct {
	Access  string
	Refresh string
}

// User defines minimal information for an authenticated user.
type User struct {
	id       entity.ID
	username string
}

// NewUser creates a new User.
func NewUser(id entity.ID, username string) *User {
	return &User{
		id:       id,
		username: username,
	}
}

// ID returns the User ID.
func (u *User) ID() entity.ID {
	return u.id
}

// Username returns the User username.
func (u *User) Username() string {
	return u.username
}
